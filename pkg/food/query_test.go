package food

import (
	"fmt"
	"github.com/artback/edamamWrapper/pkg/edamam"
	"testing"
)

func TestQuery_GetURL(t *testing.T) {
	tests := []struct {
		name    string
		fields  Query
		want    string
		wantErr bool
	}{
		{
			fields: Query{Upc: "cheese", Id: "123412421"},
			want:   fmt.Sprintf("%s?app_id=123412421&upc=cheese", parserUrl),
		},
		{
			fields: Query{Calories: edamam.Range{To: 500, From: 100}, Ingredients: []string{"cheese", "fromage"}},
			want:   fmt.Sprintf("%s?calories=100-500&ingredients=cheese&ingredients=fromage", parserUrl),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := Query{
				Key:           tt.fields.Key,
				Id:            tt.fields.Id,
				Upc:           tt.fields.Upc,
				Ingredients:   tt.fields.Ingredients,
				Health:        tt.fields.Health,
				Calories:      tt.fields.Calories,
				Category:      tt.fields.Category,
				CategoryLabel: tt.fields.CategoryLabel,
			}
			got, err := q.getURL()
			if (err != nil) != tt.wantErr {
				t.Errorf("getUrl() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getUrl() got = %v, want %v", got, tt.want)
			}
		})
	}
}
