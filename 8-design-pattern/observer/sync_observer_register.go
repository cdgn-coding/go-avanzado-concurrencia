package main

type SyncRegisterTopic struct {
	SyncTopic[string]
}

func NewSyncRegisterTopic() *SyncRegisterTopic {
	return &SyncRegisterTopic{
		SyncTopic[string]{
			observers: make([]Observer[string], 0),
		},
	}
}
