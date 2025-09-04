package decrypt

import (
	"bufio"
	"crypto/rsa"
	"errors"
	"fmt"
	"os"
	"rsa/internal/common"
	"strings"
)

type RSADecryptBatch struct {
	dataFile       string
	privateKeyFile string
	privateKey     *rsa.PrivateKey
	numbThreads    int
	messages       []string
	messagesCount  int
	doneCh         chan struct{}
	jobs           chan incomingData
	results        chan result
}

type result struct {
	decryptedMessage string
	index            int
	err              error
}

type incomingData struct {
	encryptedMessage string
	index            int
}

func NewBatch(dataFile, privateKeyFile string, numbThreads int) *RSADecryptBatch {

	return &RSADecryptBatch{
		dataFile:       dataFile,
		privateKeyFile: privateKeyFile,
		numbThreads:    numbThreads,
	}

}

func (d *RSADecryptBatch) Decrypt() (string, error) {

	var err error
	d.privateKey, err = ReadPrivateKey(d.privateKeyFile)
	if err != nil {
		return "", common.WrapError(err)
	}

	err = d.readMessages()
	if err != nil {
		return "", common.WrapError(err)
	}

	d.doneCh = make(chan struct{})
	d.jobs = d.generator(d.doneCh)
	d.results = make(chan result, d.messagesCount)

	for w := 1; w <= d.numbThreads; w++ {
		go d.worker()
	}

	return d.resultHandling()
}

func (d *RSADecryptBatch) worker() {

	for job := range d.jobs {

		select {
		case <-d.doneCh:
			return
		default:

			decryptedMessage, err := NewOnce(
				job.encryptedMessage,
				d.privateKey).Decrypt()

			d.results <- result{decryptedMessage: decryptedMessage, index: job.index, err: err}
		}
	}

}

func (d *RSADecryptBatch) generator(doneCh chan struct{}) chan incomingData {

	jobs := make(chan incomingData, len(d.messages))

	go func() {
		defer close(jobs)

		for i, data := range d.messages {
			select {
			case <-doneCh:
				return
			case jobs <- incomingData{encryptedMessage: data, index: i}:
			}
		}
	}()

	return jobs
}

func (d *RSADecryptBatch) resultHandling() (string, error) {

	resultData := make([]result, 0, d.messagesCount)

	var result result
	var errs []error
	for a := 1; a <= d.messagesCount; a++ {
		select {
		case <-d.doneCh:
			return "", common.WrapError(fmt.Errorf("operation aborted %w", errors.Join(errs...)))
		case result = <-d.results:
			if result.err != nil {
				errs = append(errs, result.err)
			}

			resultData = append(resultData, result)

		}
	}

	close(d.results)
	sortedResultData := qsort(resultData)

	stringsData := make([]string, 0, d.messagesCount)
	for _, data := range sortedResultData {
		stringsData = append(stringsData, data.decryptedMessage)
	}
	return strings.Join(stringsData, "\n"), errors.Join(errs...)

}

func (d *RSADecryptBatch) readMessages() error {

	d.messages = []string{}
	file, err := os.Open(d.dataFile)
	if err != nil {
		err = fmt.Errorf("path: %s, error: %w", d.dataFile, err)
		return common.WrapError(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		d.messages = append(d.messages, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return common.WrapError(err)
	}

	d.messagesCount = len(d.messages)
	return nil

}

func qsort(a []result) []result {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

	pivotIndex := right
	a[pivotIndex], a[right] = a[right], a[pivotIndex]

	for i := range a {
		if a[i].index < a[right].index {
			a[i], a[left] = a[left], a[i]
			left++
		}
	}

	a[left], a[right] = a[right], a[left]

	qsort(a[:left])
	qsort(a[left+1:])

	return a
}
