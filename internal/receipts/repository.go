package receipts

import (
	"github.com/google/uuid"
	cmap "github.com/orcaman/concurrent-map/v2"
)

type CMapRepo struct {
	uuidGen       UUIDGen
	concurrentMap cmap.ConcurrentMap[string, *ReceiptWithPoints]
}

func NewCMapRepo(uuidGen UUIDGen, concurrentMap cmap.ConcurrentMap[string, *ReceiptWithPoints]) *CMapRepo {
	return &CMapRepo{
		uuidGen:       uuidGen,
		concurrentMap: concurrentMap,
	}
}

func (c *CMapRepo) Save(input *ReceiptWithPoints) error {
	id := c.uuidGen()
	input.Id = id
	c.concurrentMap.Set(id.String(), input)
	return nil
}

func (c *CMapRepo) FindById(id uuid.UUID) (*ReceiptWithPoints, error) {
	found, exists := c.concurrentMap.Get(id.String())
	if !exists {
		return nil, ErrReceiptNotFound
	}
	return found, nil
}
