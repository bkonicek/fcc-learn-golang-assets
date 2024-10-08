package main

type notification interface {
	importance() int
}

type directMessage struct {
	senderUsername string
	messageContent string
	priorityLevel  int
	isUrgent       bool
}

func (dM directMessage) importance() int {
	importance := dM.priorityLevel
	if dM.isUrgent {
		importance = 50
	}
	return importance
}

type groupMessage struct {
	groupName      string
	messageContent string
	priorityLevel  int
}

func (gM groupMessage) importance() int {
	return gM.priorityLevel
}

type systemAlert struct {
	alertCode      string
	messageContent string
}

func (sA systemAlert) importance() int {
	return 100
}

func processNotification(n notification) (string, int) {
	switch v := n.(type) {
	case directMessage:
		return v.senderUsername, v.importance()
	case groupMessage:
		return v.groupName, v.importance()
	case systemAlert:
		return v.alertCode, v.importance()
	default:
		return "", 0
	}
}
