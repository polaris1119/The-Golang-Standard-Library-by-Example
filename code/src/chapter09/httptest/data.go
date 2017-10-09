package main

import (
	"errors"
	"time"
)

// 保存 Topic，没有考虑并发问题
var TopicCache = make([]*Topic, 0, 16)

type Topic struct {
	Id        int       `json:"id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

func FindTopic(id int) (*Topic, error) {
	if err := checkIndex(id); err != nil {
		return nil, err
	}

	return TopicCache[id-1], nil
}

func (t *Topic) Create() error {
	t.Id = len(TopicCache) + 1
	t.CreatedAt = time.Now()
	TopicCache = append(TopicCache, t)
	return nil
}

func (t *Topic) Update() error {
	if err := checkIndex(t.Id); err != nil {
		return err
	}
	TopicCache[t.Id-1] = t
	return nil
}

// 简单的将对应的 slice 位置置为 nil
func (t *Topic) Delete() error {
	if err := checkIndex(t.Id); err != nil {
		return err
	}
	TopicCache[t.Id-1] = nil
	return nil
}

func checkIndex(id int) error {
	if id > 0 && len(TopicCache) <= id-1 {
		return errors.New("The topic is not exists!")
	}

	return nil
}
