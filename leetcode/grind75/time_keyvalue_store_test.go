package grind75

import (
	"reflect"
	"testing"
)

// TDD => Test Driven Development
func TestTimeMapSet(t *testing.T) {
	testCases := map[string]struct {
		key    string
		values []string
		ts     []int
		want   map[string][]timeValue
	}{
		"base case": {
			key:    "foo",
			values: []string{"v1", "v2", "v3", "v4"},
			ts:     []int{1, 2, 3, 4},
			want: map[string][]timeValue{
				"foo": {
					{1, "v1"},
					{2, "v2"},
					{3, "v3"},
					{4, "v4"},
				},
			},
		},
		"repeat timestamps": {
			key:    "foo",
			values: []string{"v1", "v2", "v3", "v4", "v5", "v6"},
			ts:     []int{1, 2, 3, 4, 4, 4},
			want: map[string][]timeValue{
				"foo": {
					{1, "v1"},
					{2, "v2"},
					{3, "v3"},
					{4, "v6"},
				},
			},
		},
		"insert timestamp without order": {
			key:    "foo",
			values: []string{"v1", "v10", "v5", "v2", "v7", "v9"},
			ts:     []int{1, 10, 5, 2, 7, 9},
			want: map[string][]timeValue{
				"foo": {
					{1, "v1"},
					{2, "v2"},
					{5, "v5"},
					{7, "v7"},
					{9, "v9"},
					{10, "v10"},
				},
			},
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			storage := Constructor()
			for i := 0; i < len(tc.values); i++ {
				storage.Set(tc.key, tc.values[i], tc.ts[i])
			}
			if got := storage.KeyStore; !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}

func TestTimeMapGet(t *testing.T) {
	testCases := map[string]struct {
		keyStore  map[string][]timeValue
		key       string
		timestamp int
		want      string
	}{
		"empty key store": {
			keyStore:  map[string][]timeValue{},
			key:       "foo",
			timestamp: 1,
			want:      "",
		},
		"key is not found": {
			keyStore: map[string][]timeValue{
				"foo": {
					{1, "v1"},
					{2, "v2"},
					{3, "v3"},
					{4, "v4"},
				},
			},
			key:       "bar",
			timestamp: 1,
			want:      "",
		},
		"timestamp found in the begging": {
			keyStore: map[string][]timeValue{
				"foo": {
					{1, "v1"},
					{2, "v2"},
					{3, "v3"},
					{4, "v4"},
				},
			},
			key:       "foo",
			timestamp: 1,
			want:      "v1",
		},
		"timestamp found in the middle": {
			keyStore: map[string][]timeValue{
				"foo": {
					{1, "v1"},
					{20, "v2"},
					{35, "v3"},
					{49, "v4"},
				},
			},
			key:       "foo",
			timestamp: 35,
			want:      "v3",
		},
		"timestamp not found": {
			keyStore: map[string][]timeValue{
				"foo": {
					{1, "v1"},
					{2, "v2"},
					{3, "v3"},
					{4, "v4"},
				},
			},
			key:       "foo",
			timestamp: 10,
			want:      "v4",
		},
		"return the previous value if timestamp is not found ": {
			keyStore: map[string][]timeValue{
				"foo": {
					{1, "v1"},
					{3, "v2"},
					{5, "v3"},
					{7, "v4"},
				},
			},
			key:       "foo",
			timestamp: 4,
			want:      "v2",
		},
		"timestamp is less that any value": {
			keyStore: map[string][]timeValue{
				"love": {
					{10, "high"},
					{20, "low"},
				},
			},
			key:       "love",
			timestamp: 5,
			want:      "",
		},
		"return prev value": {
			keyStore: map[string][]timeValue{
				"love": {
					{10, "high"},
					{20, "low"},
				},
			},
			key:       "love",
			timestamp: 15,
			want:      "high",
		},
		"multiple keys": {
			keyStore: map[string][]timeValue{
				"ljfvbut": {
					{3, "tatthnvvid"},
				},
				"eimdon": {
					{8, "pdjbnnvje"},
				},
			},
			key:       "eimdon",
			timestamp: 10,
			want:      "pdjbnnvje",
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			storage := Constructor()
			storage.KeyStore = tc.keyStore
			if got := storage.Get(tc.key, tc.timestamp); !reflect.DeepEqual(got, tc.want) {
				t.Fatalf("got %v; want %v", got, tc.want)
			}
		})
	}
}
