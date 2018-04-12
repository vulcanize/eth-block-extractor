package ipfs

import "fmt"

type IpfsError struct {
	msg string
	err error
}

func (ie IpfsError) Error() string {
	return fmt.Sprintf("%s: %s", ie.msg, ie.err.Error())
}

type Publisher interface {
	Write(blockData []byte) (string, error)
}

type IpfsPublisher struct {
	DagPutter
}

func NewIpfsPublisher(dagPutter DagPutter) *IpfsPublisher {
	return &IpfsPublisher{DagPutter: dagPutter}
}

func (ip *IpfsPublisher) Write(blockData []byte) (string, error) {
	output, err := ip.DagPutter.DagPut(blockData)
	if err != nil {
		return "", err
	}
	return output, nil
}
