package coap

import (
    "os"
    "fmt"
    "reflect"
)


type oaItem struct {    // option, argument item
    Short   string      // short name
    Long    string      // long name
    Must    bool        // must exists
    HasDef  bool        // has default
    HelpMsg string      // help message
    val     reflect.Value
    rsf     reflect.StructField
}


func (oa *oaItem)parse(args []string) (cnt int) {
    // parse option/argument, return how many args used by this item
    return
}


func (oa *oaItem)init(rsf reflect.StructField, val reflect.Value) {
    oa.rsf = rsf
    oa.val = val
}


type GRP struct {       // group
    sel     string
    val     string
}


type COAP struct {
    items   []oaItem
}


/*
func (c *COAP)init() {
    if c.items != nil { return }
    //panic("init")
    c.items = make([]oaItem, 0, 10)
    st := reflect.TypeOf(*c)
    fmt.Println(st)
    fmt.Println(st.Field(0))
    fmt.Println(st.Field(1))
    fmt.Println(st.Field(2))
    fmt.Println(st.NumField())
}

func (c *COAP)Parse() { c.ParseArgs(os.Args[1:]) }
func (c *COAP)ParseArgs(args []string) {
    c.init()
}


func (c *COAP)Help() { c.HelpMsg("") }
func (c *COAP)HelpMsg(msg string) {
    c.init()
}
    //p := reflect.TypeOf(i)
    //v := reflect.ValueOf(i)
    //q := reflect.Indirect(v)
    //fmt.Println(p, q, v.CanSet(), q.CanSet())

    //ui := v.InterfaceData()
    //fmt.Println("ui =", ui)
    //fmt.Printf("ui = %d\n", v)
*/


func verifySP(i interface{}) {  // Struct Pointer
    v := reflect.ValueOf(i)
    fmt.Println("v =", v)
    k := v.Kind()
    fmt.Println("k =", k)
    if k != reflect.Ptr {
        fmt.Fprintf(os.Stderr, "Need to be a ptr\n")
        os.Exit(1)
    }
    s := reflect.Indirect(v)
    fmt.Println("s =", s)
    k = s.Kind()
    fmt.Println("k =", k)
    if k != reflect.Struct {
        fmt.Fprintf(os.Stderr, "Need to be a struct\n")
        os.Exit(1)
    }
    a := s.Addr()
    fmt.Printf("a = %d\n", a)
}


//func oneField(sf *reflect.StructField) {
//    fmt.Println("name =", sf.Name)
//    fmt.Println("tag =", sf.Tag)
//}

func initial(i interface{}) {
    verifySP(i)
    ii := reflect.Indirect(reflect.ValueOf(i))
    fmt.Println("ii =", ii)
    st := ii.Type()
    fmt.Println("st =", st)
    for idx := 0; idx < st.NumField(); idx++ {
        fs := st.Field(idx)
        //oneField(&f)
        fv := ii.Field(idx)
        fmt.Println("fv =", fv, reflect.TypeOf(fv))
        fv.SetString("MyName")
        it := &oaItem{}
        it.init(fs, fv)
    }
}


func Parse(arg interface{}) { ParseArg(arg, os.Args[1:]) }
func ParseArg(i interface{}, args []string) {
    initial(i)
    //fmt.Println("Parse(arg interface{})", arg)
    //st := reflect.TypeOf(arg)
    //fmt.Println(st)
    //fmt.Println(st.NumField())
}


func Help(arg interface{}) { HelpMsg(arg, "") }
func HelpMsg(i interface{}, msg string) {
    initial(i)
}
