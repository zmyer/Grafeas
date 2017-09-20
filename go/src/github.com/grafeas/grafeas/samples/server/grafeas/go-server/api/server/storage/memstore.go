package storage

import (
	"fmt"
	"github.com/grafeas/grafeas/samples/server/grafeas/go-server/api"
	"github.com/grafeas/grafeas/samples/server/grafeas/go-server/api/server/errors"

	"net/http"
)

type Store interface {
	CreateOccurrence(o *swagger.Occurrence) *errors.AppError
	DeleteOccurrence(projectID, oID string) *errors.AppError
	UpdateOccurrence(projectID, oID string, o *swagger.Occurrence) *errors.AppError
	GetOccurrence(projectID, oID string) (*swagger.Occurrence, *errors.AppError)
	ListOccurrences() *errors.AppError
	CreateNote(n *swagger.Note) *errors.AppError
	DeleteNote(providerID, nID string) *errors.AppError
	UpdateNote(providerID, nID string, n *swagger.Note) *errors.AppError
	GetNote(providerID, nID string, n *swagger.Note) *errors.AppError
	GetNoteByOccurrence(projectID, oID string) (*swagger.Note, *errors.AppError)
	ListNotes() *errors.AppError
	ListNoteOccurrences() *errors.AppError

	GetOperation(projectID, opID string) (*swagger.Operation, *errors.AppError)
	CreateOperation(o *swagger.Operation) *errors.AppError
	DeleteOperation(projectID, opID string) *errors.AppError
	UpdateOperation(projectID, opID string, op *swagger.Operation) *errors.AppError
	ListOperations() *errors.AppError
}

type MemStore struct {
	occurrencesByID map[string]swagger.Occurrence
	notesByID       map[string]swagger.Note
	opsByID         map[string]swagger.Operation
}

func NewMemStore() *MemStore {
	return &MemStore{make(map[string]swagger.Occurrence), make(map[string]swagger.Note),
		make(map[string]swagger.Operation)}
}

func occurrenceName(pID, oID string) string {
	return fmt.Sprintf("projects/%v/occurrences/%v", pID, oID)
}

func noteName(pID, nID string) string {
	return fmt.Sprintf("projects/%v/notes/%v", pID, nID)
}

func (m *MemStore) CreateOccurrence(o *swagger.Occurrence) *errors.AppError {
	if _, ok := m.occurrencesByID[o.Name]; ok {
		return &errors.AppError{fmt.Sprintf("Occurrence with Name %v already exists", o.Name),
			http.StatusBadRequest}
	}
	m.occurrencesByID[o.Name] = *o
	return nil
}

func (m *MemStore) DeleteOccurrence(pID, oID string) *errors.AppError {
	name := occurrenceName(pID, oID)
	if _, ok := m.occurrencesByID[name]; !ok {
		return &errors.AppError{fmt.Sprintf("Occurrence with Name %v does not Exist", name),
			http.StatusBadRequest}
	}
	delete(m.occurrencesByID, name)
	return nil
}

func (m *MemStore) UpdateOccurrence(projectID, oID string, o *swagger.Occurrence) *errors.AppError {
	//name := fmt.Sprintf("projects/%v/occurrences/%v", projectID, oID)
	panic("implement me")

}

func (m *MemStore) GetOccurrence(projectID, oID string) (*swagger.Occurrence, *errors.AppError) {
	panic("implement me")
}

func (m *MemStore) ListOccurrences() *errors.AppError {
	panic("implement me")
}

func (m *MemStore) CreateNote(n *swagger.Note) *errors.AppError {
	if _, ok := m.notesByID[n.Name]; ok {
		return &errors.AppError{fmt.Sprintf("Occurrence with Name %v already exists", n.Name),
			http.StatusBadRequest}
	}
	m.notesByID[n.Name] = *n
	return nil
}

func (m *MemStore) DeleteNote(providerID, nID string) *errors.AppError {
	panic("implement me")
}

func (m *MemStore) UpdateNote(providerID, nID string, n *swagger.Note) *errors.AppError {
	panic("implement me")
}

func (m *MemStore) GetNote(providerID, nID string, n *swagger.Note) *errors.AppError {
	panic("implement me")
}

func (m *MemStore) GetNoteByOccurrence(projectID, oID string) (*swagger.Note, *errors.AppError) {
	panic("implement me")
}

func (m *MemStore) ListNotes() *errors.AppError {
	panic("implement me")
}

func (m *MemStore) ListNoteOccurrences() *errors.AppError {
	panic("implement me")
}

func (m *MemStore) GetOperation(projectID, opID string) (*swagger.Operation, *errors.AppError) {
	panic("implement me")
}

func (m *MemStore) CreateOperation(o *swagger.Operation) *errors.AppError {
	if _, ok := m.opsByID[o.Name]; ok {
		return &errors.AppError{fmt.Sprintf("Operation with Name %v already exists", o.Name),
			http.StatusBadRequest}
	}
	m.opsByID[o.Name] = *o
	return nil
}

func (m *MemStore) DeleteOperation(projectID, opID string) *errors.AppError {
	panic("implement me")
}

func (m *MemStore) UpdateOperation(projectID, opID string, op *swagger.Operation) *errors.AppError {
	panic("implement me")
}

func (m *MemStore) ListOperations() *errors.AppError {
	panic("implement me")
}
