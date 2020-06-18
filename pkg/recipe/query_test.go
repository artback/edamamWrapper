package recipe

import (
	"fmt"
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
			fields: Query{SearchText: "cheese", Id: "123412421"},
			want:   fmt.Sprintf("%s?app_id=123412421&q=cheese", searchUrl),
		},
		{
			fields: Query{MaxIngredients: 5, Excluded: []string{"cheese", "fromage"}},
			want:   fmt.Sprintf("%s?excluded=cheese&excluded=fromage&ingr=5", searchUrl),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			q := Query{
				SearchText:     tt.fields.SearchText,
				Key:            tt.fields.Key,
				Id:             tt.fields.Id,
				MaxIngredients: tt.fields.MaxIngredients,
				DietLabel:      tt.fields.DietLabel,
				HealthLabel:    tt.fields.HealthLabel,
				CuisineType:    tt.fields.CuisineType,
				MealType:       tt.fields.MealType,
				DishType:       tt.fields.DishType,
				Calories:       tt.fields.Calories,
				Time:           tt.fields.Time,
				Excluded:       tt.fields.Excluded,
			}
			got, err := q.getUrl()
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
