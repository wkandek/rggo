package todo_test

import (
    "io/ioutil"
    "os"
    "testing"
    "github.com/wkandek/rggo/todo"
)


func TestAdd(t *testing.T) {
    l := todo.List{}
    taskName := "New task"
    l.Add(taskName)
    if l[0].Task != taskName {
        t.Errorf("expect %q, got %q instaed.", taskName, l[0].Task)
    }
}

func TestComplete( t *testing.T) {
    l := todo.List{}
    taskName := "New task"
    l.Add(taskName)
    if l[0].Task != taskName {
        t.Errorf("expect %q, got %q instaed.", taskName, l[0].Task)
    }
    if l[0].Done {
        t.Errorf("Task should not be done")
    }
    l.Complete(1)
    if !l[0].Done {
        t.Errorf("Task should be done")
    }
}

func TestDelete(t *testing.T) {
    l := todo.List{}
    tasks := []string{
        "task 1",
        "task 2",
        "task 3",
    }
    for _, v := range tasks {
        l.Add(v)
    }

    for i := 0; i < len(l); i++ {
        if l[i].Task != tasks[i] {
            t.Errorf("expect %q, got %q %d", l[i].Task, tasks[i], i)
        }
    }

    l.Delete(2)

    if l[1].Task != tasks[2] {
        t.Errorf("After Delete got %q, expected %q, %d", l[1].Task, tasks[2], len(l))
    }
}

func TestSaveGet(t *testing.T){
    l1 := todo.List{}
    l2 := todo.List{}
    taskName := "New task"
    l1.Add(taskName)
    if l1[0].Task != taskName {
        t.Errorf("expect %q, got %q instaed.", taskName, l1[0].Task)
    }
    tf, err := ioutil.TempFile("", "")
    if err != nil {
        t.Fatalf("Error creating tem file %s", err)
    }
    defer os.Remove(tf.Name())
    if err := l1.Save(tf.Name()); err != nil {
        t.Fatalf("Error saving list to file: %s", err)
    }
    if err := l2.Get(tf.Name()); err != nil {
        t.Fatalf("Error getting list to file: %s", err)
    }
    if l1[0].Task != l2[0].Task {
        t.Errorf(" Task %q should match %q", l1[0].Task, l2[0].Task)
    }
}
