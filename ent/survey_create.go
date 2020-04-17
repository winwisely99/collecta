// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/facebookincubator/ent/schema/field"
	"github.com/google/uuid"
	"github.com/minskylab/collecta/ent/domain"
	"github.com/minskylab/collecta/ent/flow"
	"github.com/minskylab/collecta/ent/survey"
	"github.com/minskylab/collecta/ent/user"
)

// SurveyCreate is the builder for creating a Survey entity.
type SurveyCreate struct {
	config
	id              *uuid.UUID
	tags            *[]string
	lastInteraction *time.Time
	dueDate         *time.Time
	title           *string
	description     *string
	metadata        *map[string]string
	done            *bool
	flow            map[uuid.UUID]struct{}
	_for            map[uuid.UUID]struct{}
	owner           map[uuid.UUID]struct{}
}

// SetTags sets the tags field.
func (sc *SurveyCreate) SetTags(s []string) *SurveyCreate {
	sc.tags = &s
	return sc
}

// SetLastInteraction sets the lastInteraction field.
func (sc *SurveyCreate) SetLastInteraction(t time.Time) *SurveyCreate {
	sc.lastInteraction = &t
	return sc
}

// SetDueDate sets the dueDate field.
func (sc *SurveyCreate) SetDueDate(t time.Time) *SurveyCreate {
	sc.dueDate = &t
	return sc
}

// SetNillableDueDate sets the dueDate field if the given value is not nil.
func (sc *SurveyCreate) SetNillableDueDate(t *time.Time) *SurveyCreate {
	if t != nil {
		sc.SetDueDate(*t)
	}
	return sc
}

// SetTitle sets the title field.
func (sc *SurveyCreate) SetTitle(s string) *SurveyCreate {
	sc.title = &s
	return sc
}

// SetDescription sets the description field.
func (sc *SurveyCreate) SetDescription(s string) *SurveyCreate {
	sc.description = &s
	return sc
}

// SetNillableDescription sets the description field if the given value is not nil.
func (sc *SurveyCreate) SetNillableDescription(s *string) *SurveyCreate {
	if s != nil {
		sc.SetDescription(*s)
	}
	return sc
}

// SetMetadata sets the metadata field.
func (sc *SurveyCreate) SetMetadata(m map[string]string) *SurveyCreate {
	sc.metadata = &m
	return sc
}

// SetDone sets the done field.
func (sc *SurveyCreate) SetDone(b bool) *SurveyCreate {
	sc.done = &b
	return sc
}

// SetNillableDone sets the done field if the given value is not nil.
func (sc *SurveyCreate) SetNillableDone(b *bool) *SurveyCreate {
	if b != nil {
		sc.SetDone(*b)
	}
	return sc
}

// SetID sets the id field.
func (sc *SurveyCreate) SetID(u uuid.UUID) *SurveyCreate {
	sc.id = &u
	return sc
}

// SetFlowID sets the flow edge to Flow by id.
func (sc *SurveyCreate) SetFlowID(id uuid.UUID) *SurveyCreate {
	if sc.flow == nil {
		sc.flow = make(map[uuid.UUID]struct{})
	}
	sc.flow[id] = struct{}{}
	return sc
}

// SetFlow sets the flow edge to Flow.
func (sc *SurveyCreate) SetFlow(f *Flow) *SurveyCreate {
	return sc.SetFlowID(f.ID)
}

// SetForID sets the for edge to User by id.
func (sc *SurveyCreate) SetForID(id uuid.UUID) *SurveyCreate {
	if sc._for == nil {
		sc._for = make(map[uuid.UUID]struct{})
	}
	sc._for[id] = struct{}{}
	return sc
}

// SetFor sets the for edge to User.
func (sc *SurveyCreate) SetFor(u *User) *SurveyCreate {
	return sc.SetForID(u.ID)
}

// SetOwnerID sets the owner edge to Domain by id.
func (sc *SurveyCreate) SetOwnerID(id uuid.UUID) *SurveyCreate {
	if sc.owner == nil {
		sc.owner = make(map[uuid.UUID]struct{})
	}
	sc.owner[id] = struct{}{}
	return sc
}

// SetNillableOwnerID sets the owner edge to Domain by id if the given value is not nil.
func (sc *SurveyCreate) SetNillableOwnerID(id *uuid.UUID) *SurveyCreate {
	if id != nil {
		sc = sc.SetOwnerID(*id)
	}
	return sc
}

// SetOwner sets the owner edge to Domain.
func (sc *SurveyCreate) SetOwner(d *Domain) *SurveyCreate {
	return sc.SetOwnerID(d.ID)
}

// Save creates the Survey in the database.
func (sc *SurveyCreate) Save(ctx context.Context) (*Survey, error) {
	if sc.tags == nil {
		return nil, errors.New("ent: missing required field \"tags\"")
	}
	if sc.lastInteraction == nil {
		return nil, errors.New("ent: missing required field \"lastInteraction\"")
	}
	if sc.dueDate == nil {
		v := survey.DefaultDueDate()
		sc.dueDate = &v
	}
	if sc.title == nil {
		return nil, errors.New("ent: missing required field \"title\"")
	}
	if err := survey.TitleValidator(*sc.title); err != nil {
		return nil, fmt.Errorf("ent: validator failed for field \"title\": %v", err)
	}
	if sc.done == nil {
		v := survey.DefaultDone
		sc.done = &v
	}
	if len(sc.flow) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"flow\"")
	}
	if sc.flow == nil {
		return nil, errors.New("ent: missing required edge \"flow\"")
	}
	if len(sc._for) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"for\"")
	}
	if sc._for == nil {
		return nil, errors.New("ent: missing required edge \"for\"")
	}
	if len(sc.owner) > 1 {
		return nil, errors.New("ent: multiple assignments on a unique edge \"owner\"")
	}
	return sc.sqlSave(ctx)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SurveyCreate) SaveX(ctx context.Context) *Survey {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (sc *SurveyCreate) sqlSave(ctx context.Context) (*Survey, error) {
	var (
		s     = &Survey{config: sc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: survey.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeUUID,
				Column: survey.FieldID,
			},
		}
	)
	if value := sc.id; value != nil {
		s.ID = *value
		_spec.ID.Value = *value
	}
	if value := sc.tags; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  *value,
			Column: survey.FieldTags,
		})
		s.Tags = *value
	}
	if value := sc.lastInteraction; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: survey.FieldLastInteraction,
		})
		s.LastInteraction = *value
	}
	if value := sc.dueDate; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  *value,
			Column: survey.FieldDueDate,
		})
		s.DueDate = *value
	}
	if value := sc.title; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: survey.FieldTitle,
		})
		s.Title = *value
	}
	if value := sc.description; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  *value,
			Column: survey.FieldDescription,
		})
		s.Description = *value
	}
	if value := sc.metadata; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  *value,
			Column: survey.FieldMetadata,
		})
		s.Metadata = *value
	}
	if value := sc.done; value != nil {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  *value,
			Column: survey.FieldDone,
		})
		s.Done = *value
	}
	if nodes := sc.flow; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   survey.FlowTable,
			Columns: []string{survey.FlowColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: flow.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc._for; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   survey.ForTable,
			Columns: []string{survey.ForColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: user.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := sc.owner; len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   survey.OwnerTable,
			Columns: []string{survey.OwnerColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeUUID,
					Column: domain.FieldID,
				},
			},
		}
		for k, _ := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
		if cerr, ok := isSQLConstraintError(err); ok {
			err = cerr
		}
		return nil, err
	}
	return s, nil
}
