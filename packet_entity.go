package manta

import (
	"github.com/dotabuff/manta/dota"
)

// Represents the state of an entity
type packetEntity struct {
	index      int32
	classId    int32
	className  string
	flatTbl    *dt
	properties map[string]interface{}
}

// Internal callback for CSVCMsg_PacketEntities.
func (p *Parser) onCSVCMsg_PacketEntities(m *dota.CSVCMsg_PacketEntities) error {
	// XXX: Remove once we've gotten ReadProperties working.
	return nil

	/*defer func() {
		if err := recover(); err != nil {
			_debugf("recovered: %s", err)
		}
	}()*/

	_debugf("pTick=%d isDelta=%v deltaFrom=%d updatedEntries=%d maxEntries=%d baseline=%d updateBaseline=%v", p.Tick, m.GetIsDelta(), m.GetDeltaFrom(), m.GetUpdatedEntries(), m.GetMaxEntries(), m.GetBaseline(), m.GetUpdateBaseline())

	r := NewReader(m.GetEntityData())
	index := int32(-1)
	ok := false

	// Iterate over all entries
	for i := 0; i < int(m.GetUpdatedEntries()); i++ {
		// Read the index delta from the buffer. This is an implementation
		// from Alice. An alternate implementation from Yasha has the same result.
		delta := r.readUBitVar()
		index += int32(delta) + 1
		_debugf("index delta is %d to %d", delta, index)

		// Read the type of update based on two booleans.
		// This appears to be backwards from source 1:
		// true+true used to be "create", now appears to be false+true?
		// This seems suspcious.
		updateType := " "
		if r.readBoolean() {
			if r.readBoolean() {
				updateType = "D"
			} else {
				updateType = "?"
			}
		} else {
			if r.readBoolean() {
				updateType = "C"
			} else {
				updateType = "U"
			}
		}
		_debugf("update type is %s", updateType)

		// Proceed based on the update type
		switch updateType {
		case "C":
			// Create a new packetEntity.
			pe := &packetEntity{
				index:      index,
				classId:    int32(r.readBits(p.classIdSize)),
				properties: make(map[string]interface{}),
			}

			// Skip the 10 serial bits for now.
			r.seekBits(10)

			// Get the associated class.
			if pe.className, ok = p.ClassInfo[pe.classId]; !ok {
				_panicf("unable to find class %d", pe.classId)
			}

			// Get the associated serializer
			if pe.flatTbl, ok = p.Serializers[pe.className][0]; !ok {
				_panicf("unable to find serializer for class %s", pe.className)
			}

			// Register the packetEntity with the parser.
			p.packetEntities[index] = pe

			_debugf("created a pe: %+v", pe)

			// Read properties and set them in the packetEntity
			pe.properties = ReadPropertiesNew(r, pe.flatTbl)
		case "U":
			_panicf("End here")
			// Find the existing packetEntity
			pe, ok := p.packetEntities[index]
			if !ok {
				_debugf("unable to find packet entity %d for update", index)
			}

			return nil

			// Read properties and update the packetEntity
			for k, v := range ReadPropertiesNew(r, pe.flatTbl) {
				pe.properties[k] = v
			}

		case "D":
			return nil
			if _, ok := p.packetEntities[index]; !ok {
				_panicf("unable to find packet entity %d for delete", index)
			} else {
				delete(p.packetEntities, index)
			}
		}
	}

	_panicf("End here")

	return nil
}
