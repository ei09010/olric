package dmap

type MemCacheMap struct {
	DMap *DMap
}

func (s *Service) NewMemCacheMap(name string) (*MemCacheMap, error) {

	dmap, err := s.NewDMap(name)

	if err != nil {
		return nil, err
	}

	mcMap := &MemCacheMap{
		DMap: dmap,
	}

	return mcMap, nil
}
