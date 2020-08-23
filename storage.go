package spruce

// 用于持久化存储

func (h *Hash) Storage() {
	FindAll(h.ver)
}
