package models

type Computer struct {
	ID            int    `json:"id"`
	ProcessorCore int    `json:"processor_core"`
	StorageSize   int    `json:"storage_size"`
	Memory        int    `json:"memory"`
	VideoCard     string `json:"video_card"`
}
