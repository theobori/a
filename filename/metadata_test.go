package filename

import (
	"fmt"
	"testing"
)

func TestNewMetadata(t *testing.T) {
	testsOK := []struct {
		title            string
		tags             []string
		ext              string
		expectedFilename string
	}{
		{title: "test1", tags: []string{"tag1", "tag2"}, expectedFilename: "test1__tag1_tag2", ext: ""},
		{title: "test2", tags: []string{"tag1", "tag1", "tag2", "tag2"}, expectedFilename: "test2__tag1_tag2", ext: ""},
		{title: "test3", tags: []string{"  tag1  ", "tag1", "tag2      ", "tag2"}, expectedFilename: "test3__tag1_tag2", ext: ""},
		{title: "test4", tags: []string{}, expectedFilename: "test4", ext: ""},
		{title: "test5", tags: []string{}, expectedFilename: "test5", ext: ""},
		{title: "test6", tags: []string{"4", "3", "1", "2"}, expectedFilename: "test6__1_2_3_4", ext: ""},
	}

	for _, testOK := range testsOK {
		fm, err := NewMetadata(testOK.title, testOK.tags, testOK.ext)
		if err != nil {
			t.Fatal(err)
		}

		filename := fm.Filename()
		if filename != testOK.expectedFilename {
			t.Fatalf(
				"'%s' is not equal to '%s'",
				filename,
				testOK.expectedFilename,
			)
		}
	}

	testsKO := []struct {
		title string
		tags  []string
		ext   string
	}{
		{title: "hello__world", tags: []string{}, ext: ""},
		{title: "working-title", tags: []string{"tag1", "tag__2"}, ext: ""},
		{title: "", tags: []string{"tag1"}, ext: ""},
		{title: "title", tags: []string{}, ext: "zzz"},
	}

	for _, testKO := range testsKO {
		_, err := NewMetadata(testKO.title, testKO.tags, testKO.ext)
		if err == nil {
			t.Fatalf(
				"Filename with the title '%s' and the tags %v should have failed test.",
				testKO.title,
				testKO.tags,
			)
		}
	}
}

func TestNewMetadataFromPath(t *testing.T) {
	optionsDefault := NewOptionsDefault()

	testsOK := []struct {
		path       string
		tagsAmount int
	}{
		{path: "a/b/c/my_title__1_2_3.txt", tagsAmount: 3},
		{path: "/a/b/c/title__1.org", tagsAmount: 1},
		{path: "/a/b/c/title.md", tagsAmount: 0},
		{path: "/a/b/c/title", tagsAmount: 0},
		{path: "my_title", tagsAmount: 0},
	}

	for _, testOK := range testsOK {
		fm, err := NewMetadataFromPath(testOK.path, optionsDefault)
		if err != nil {
			t.Fatal(err)
		}

		tagsAmount := fm.Tags.Len()
		if tagsAmount != testOK.tagsAmount {
			t.Fatalf(
				"The tag amount of '%s' is %d and should be %d.",
				testOK.path,
				tagsAmount,
				testOK.tagsAmount,
			)
		}
	}

	testsKO := []struct {
		path string
	}{
		{path: "/a/b/c/title__"},
		{path: "/a/b/c/tit___le__a_b"},
		{path: "/a/b/c/__tag1_tag2"},
		{path: "/a/b/c/__"},
		{path: "  "},
	}

	for _, testKO := range testsKO {
		fm, err := NewMetadataFromPath(testKO.path, optionsDefault)
		if err == nil {
			fmt.Println(fm.Title, fm.Tags, fm.Ext)
			t.Fatalf("The path '%s' should have failed test.", testKO.path)
		}
	}
}
