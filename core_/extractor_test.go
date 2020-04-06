package core_

import (
	"reflect"
	"testing"
)

func TestExtractObject(t *testing.T) {
	type args struct {
		text   string
		anchor *Anchor
	}
	tests := []struct {
		name    string
		args    args
		want    *Extract
		wantErr error
	}{
		{
			name: "foundElement",
			args: args{
				text:   "Hello, how are you? Hey, What do you doing?",
				anchor: NewAnchor(", ", " are"),
			},
			want: &Extract{
				Object: "how",
				SPos:   7,
				FPos:   10,
			},
			wantErr: nil,
		},
		{
			name: "notFoundFinishAnchor",
			args: args{
				text:   "Hello, how are you? Hey, What are you doing?",
				anchor: NewAnchor(", ", " ARE"),
			},
			want:    nil,
			wantErr: BadHypertextError,
		},
		{
			name: "notFoundStartAnchor",
			args: args{
				text:   "Hello, how are you? Hey, What are you doing?",
				anchor: NewAnchor(", H", " are"),
			},
			want:    nil,
			wantErr: NotFoundObjectError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractObject(tt.args.text, tt.args.anchor)
			if (err != nil) && err != tt.wantErr {
				t.Errorf("ExtractObject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractObject() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractObjects(t *testing.T) {
	type args struct {
		text   string
		anchor *Anchor
		count  *uint32
	}
	tests := []struct {
		name    string
		args    args
		want    []*Extract
		wantErr error
	}{
		{
			name: "foundElement",
			args: args{
				text:   "Hello, how are you? Hey! What do you doing?",
				anchor: NewAnchor(", ", " are"),
			},
			want: []*Extract{
				{
					Object: "how",
					SPos:   7,
					FPos:   10,
				},
			},
			wantErr: nil,
		},
		{
			name: "foundTwoElements",
			args: args{
				text:   "Hello, how are you? Hey, what are you doing?",
				anchor: NewAnchor(", ", " are"),
			},
			want: []*Extract{
				{
					Object: "how",
					SPos:   7,
					FPos:   10,
				},
				{
					Object: "what",
					SPos:   25,
					FPos:   29,
				},
			},
			wantErr: nil,
		},
		{
			name: "notFoundFinishAnchorInTheLast",
			args: args{
				text:   "Hello, how are you? Hey, What do you doing?",
				anchor: NewAnchor(", ", " are"),
			},
			want:    nil,
			wantErr: BadHypertextError,
		},
		{
			name: "notFoundFinishAnchor",
			args: args{
				text:   "Hello, how are you? Hey, What are you doing?",
				anchor: NewAnchor(", ", " ARE"),
			},
			want:    nil,
			wantErr: BadHypertextError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractObjects(tt.args.text, tt.args.anchor, nil)
			if (err != nil) && err != tt.wantErr {
				t.Errorf("ExtractObjects() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractObjects() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TODO: this tests not implemented therefore this function may be reducing
func TestExtractCtxObjects(t *testing.T) {
	type args struct {
		text      string
		ctxAnchor *CtxAnchor
		count     *uint32
	}
	tests := []struct {
		name    string
		args    args
		want    []*Extract
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractCtxObjects(tt.args.text, tt.args.ctxAnchor, tt.args.count)
			if (err != nil) != tt.wantErr {
				t.Errorf("ExtractCtxObjects() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractCtxObjects() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExtractCtxObject(t *testing.T) {
	type args struct {
		text      string
		ctxAnchor *CtxAnchor
	}
	tests := []struct {
		name    string
		args    args
		want    *Extract
		wantErr error
	}{
		{
			name: "foundWithContext",
			args: args{
				text:      "<x>Hello world!<z><x>a</x></z>",
				ctxAnchor: NewCtxAnchor("<z>", NewAnchor("<x>", "</x>")),
			},
			want: &Extract{
				Object: "a",
				SPos:   20,
				FPos:   21,
			},
			wantErr: nil,
		},
		{
			name: "notFoundFinishAnchor",
			args: args{
				text:      "<x>Hello world!<z><x>a</x></z>",
				ctxAnchor: NewCtxAnchor("<z>", NewAnchor("<x>", "<//x>")),
			},
			want:    nil,
			wantErr: BadHypertextError,
		},
		{
			name: "notFoundStartAnchor",
			args: args{
				text:      "<x>Hello world!<z><x>a</x></z>",
				ctxAnchor: NewCtxAnchor("<z>", NewAnchor("<y>", "</x>")),
			},
			want:    nil,
			wantErr: NotFoundObjectError,
		},
		{
			name: "notFoundCtxAnchor",
			args: args{
				text:      "<x>Hello world!<z><x>a</x></z>",
				ctxAnchor: NewCtxAnchor("<y>", NewAnchor("<x>", "</x>")),
			},
			want:    nil,
			wantErr: NotFoundCtxError,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ExtractCtxObject(tt.args.text, tt.args.ctxAnchor)
			if (err != nil) && err != tt.wantErr {
				t.Errorf("ExtractCtxObject() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ExtractCtxObject() = %v, want %v", got, tt.want)
			}
		})
	}
}
