package user

import (
)

func (h *handler) FindHandler(req *request) (any, error) {
	m, err := h.userRepo.FindOneBy(map[string]interface{}{
		"id": req.ID,
	})
	if err != nil {
		return nil, err
	}

	return m, nil
}
