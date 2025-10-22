package main

import (
    "container/list"
)

type OptimalList struct {
    size int
    data interface{}
}

func NewOptimalList() *OptimalList {
    return &OptimalList{
        size: 0,
        data: nil,
    }
}

func (ol *OptimalList) Add(value interface{}) {
    ol.size++
    
    switch {
    case ol.size == 1:
        ol.data = value
        
    case ol.size >= 2 && ol.size <= 5:
        if ol.size == 2 {
            arr := [5]interface{}{ol.data, value}
            ol.data = arr
        } else {
            arr := ol.data.([5]interface{})
            arr[ol.size-1] = value
            ol.data = arr
        }
        
    case ol.size == 6:
        arr := ol.data.([5]interface{})
        l := list.New()
        for i := 0; i < 5; i++ {
            l.PushBack(arr[i])
        }
        l.PushBack(value)
        ol.data = l
        
    default:
        l := ol.data.(*list.List)
        l.PushBack(value)
    }
}

func (ol *OptimalList) Remove() interface{} {
    if ol.size == 0 {
        return nil
    }
    
    var value interface{}
    
    switch {
    case ol.size == 1:
        value = ol.data
        ol.data = nil
        
    case ol.size >= 2 && ol.size <= 5:
        arr := ol.data.([5]interface{})
        value = arr[ol.size-1]
        arr[ol.size-1] = nil
        
        if ol.size == 2 {
            ol.data = arr[0]
        } else {
            ol.data = arr
        }
        
    case ol.size == 6:
        l := ol.data.(*list.List)
        last := l.Back()
        value = last.Value
        l.Remove(last)
        
        arr := [5]interface{}{}
        i := 0
        for e := l.Front(); e != nil; e = e.Next() {
            arr[i] = e.Value
            i++
        }
        ol.data = arr
        
    default:
        l := ol.data.(*list.List)
        last := l.Back()
        value = last.Value
        l.Remove(last)
    }
    
    ol.size--
    return value
}

func (ol *OptimalList) Get(index int) interface{} {
    if index < 0 || index >= ol.size {
        return nil
    }
    
    switch {
    case ol.size == 1:
        return ol.data
        
    case ol.size >= 2 && ol.size <= 5:
        arr := ol.data.([5]interface{})
        return arr[index]
        
    default:
        l := ol.data.(*list.List)
        i := 0
        for e := l.Front(); e != nil; e = e.Next() {
            if i == index {
                return e.Value
            }
            i++
        }
        return nil
    }
}

func (ol *OptimalList) Set(index int, value interface{}) bool {
    if index < 0 || index >= ol.size {
        return false
    }
    
    switch {
    case ol.size == 1:
        ol.data = value
        return true
        
    case ol.size >= 2 && ol.size <= 5:
        arr := ol.data.([5]interface{})
        arr[index] = value
        ol.data = arr
        return true
        
    default:
        l := ol.data.(*list.List)
        i := 0
        for e := l.Front(); e != nil; e = e.Next() {
            if i == index {
                e.Value = value
                return true
            }
            i++
        }
        return false
    }
}

func (ol *OptimalList) Size() int {
    return ol.size
}

func (ol *OptimalList) String() string {
    result := "["
    
    switch {
    case ol.size == 0:
    
    case ol.size == 1:
        result += toString(ol.data)
        
    case ol.size >= 2 && ol.size <= 5:
        arr := ol.data.([5]interface{})
        for i := 0; i < ol.size; i++ {
            if i > 0 {
                result += ", "
            }
            result += toString(arr[i])
        }
        
    default:
        l := ol.data.(*list.List)
        first := true
        for e := l.Front(); e != nil; e = e.Next() {
            if !first {
                result += ", "
            }
            result += toString(e.Value)
            first = false
        }
    }
    
    result += "]"
    return result
}

func toString(v interface{}) string {
    if v == nil {
        return "nil"
    }
    return "value"
}

func main() {
    list := NewOptimalList()
    
    for i := 0; i < 10; i++ {
        list.Add(i)
        println("After adding", i, "Size:", list.Size(), "Content:", list.String())
    }
    
    for i := 0; i < list.Size(); i++ {
        println("Get(", i, ") =", list.Get(i))
    }
    
    list.Set(5, 100)
    println("After setting index 5 to 100:", list.String())
    
    for list.Size() > 0 {
        removed := list.Remove()
        println("Removed:", removed, "Size:", list.Size(), "Content:", list.String())
    }
}