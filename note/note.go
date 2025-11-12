package note

import (
	"errors"
	"regexp"
	"time"
	"util/middle"
	"util/security"
)

var notesCache middle.ExpirableCache

type note struct {
	Content     string
	ReadOnce    bool
	CreatedOn   time.Time
	TTLSeconds  int
	IsEncrypted bool
}

type NoteDto struct {
	NoteId      string    `json:"id"`
	Content     string    `json:"content"`
	ReadOnce    bool      `json:"readOnce"`
	CreatedOn   time.Time `json:"createdOn"`
	TTLSeconds  int       `json:"ttlSeconds"`
	IsEncrypted bool      `json:"isEncrypted"`
}

func validateArgs(noteDto NoteDto) error {
	if len(noteDto.NoteId) < 3 {
		return errors.New("note id should have at least three characters")
	}
	if len(noteDto.NoteId) > 100 {
		return errors.New("note id should have 100 characters at most")
	}
	if len(noteDto.Content) > 25000 {
		return errors.New("note content should have 25000 characters at most")
	}
	if noteDto.TTLSeconds > 18000 {
		return errors.New("the maximum duration is 18000 seconds")
	}
	match, _ := regexp.Match(`^[A-Za-z0-9\-]+$`, []byte(noteDto.NoteId))
	if !match {
		return errors.New("the note id should have only alphanumeric characters")
	}

	return nil
}

func CreateNote(noteDto NoteDto) error {
	_, found := notesCache.Get(noteDto.NoteId)
	if found {
		return errors.New("note already exists with this id")
	}

	err := validateArgs(noteDto)
	if err != nil {
		return err
	}

	note := note{
		Content:     security.EscapeHTML(noteDto.Content),
		ReadOnce:    noteDto.ReadOnce,
		CreatedOn:   noteDto.CreatedOn,
		TTLSeconds:  noteDto.TTLSeconds,
		IsEncrypted: noteDto.IsEncrypted,
	}

	notesCache.Set(noteDto.NoteId, note, time.Second*time.Duration(noteDto.TTLSeconds))
	return nil
}

func GetNote(noteId string) (*NoteDto, error) {
	cacheRaw, ok := notesCache.Get(noteId)
	if !ok {
		return nil, errors.New("note not found")
	}
	noteFromCache := cacheRaw.(note)
	if noteFromCache.ReadOnce {
		notesCache.Delete(noteId)
	}
	noteDto := NoteDto{
		NoteId:      noteId,
		Content:     noteFromCache.Content,
		ReadOnce:    noteFromCache.ReadOnce,
		CreatedOn:   noteFromCache.CreatedOn,
		TTLSeconds:  noteFromCache.TTLSeconds,
		IsEncrypted: noteFromCache.IsEncrypted,
	}
	return &noteDto, nil
}
