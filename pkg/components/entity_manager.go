package components

import (
	"fmt"
)

type EntityManager struct {
	entities []Entity

	Cameras   map[Entity]*Camera
	Positions map[Entity]*Position
	Momentums map[Entity]*Momentum

	entityCounter func() Entity
}

func NewEntityManager() *EntityManager {
	return &EntityManager{
		entities:      []Entity{},
		Cameras:       map[Entity]*Camera{},
		Positions:     map[Entity]*Position{},
		Momentums:     map[Entity]*Momentum{},
		entityCounter: NewIDGenerator(),
	}
}

func (m *EntityManager) NewEntity(comps ...interface{}) Entity {
	entity := m.entityCounter()
	for _, comp := range comps {
		switch c := comp.(type) {
		case *Camera:
			m.Cameras[entity] = c
			break
		case *Position:
			m.Positions[entity] = c
			break
		case *Momentum:
			m.Momentums[entity] = c
			break

		default:
			fmt.Println("Component not in define in the manager")
		}
	}
	m.entities = append(m.entities, entity)

	return entity
}
