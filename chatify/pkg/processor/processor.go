package processor

type Processor struct {
	First Handler
	End   Handler
}

func (p *Processor) Start(c *MContext) (err error) {
	curr := p.First

	for curr != nil && err == nil {
		err = curr.Process(c)
		curr = curr.Next()
	}

	return err
}

func (p *Processor) Append(h Handler) {
	if p.First == nil {
		p.First = h
		p.End = h
		return
	}

	p.End.SetNext(h)
	p.End = h
}

func InitMsgProcessor(formatter *Formatter, persistency *Persistency, customs ...Handler) *Processor {
	var processor Processor
	processor.Append(formatter)

	if persistency != nil {
		processor.Append(persistency)
	}

	for i := range customs {
		processor.Append(customs[i])
	}
	return &processor
}
