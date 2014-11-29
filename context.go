package main

import (
  "log"
  "os"
  )


type Context struct {
  arguments *Arguments
  debugLog       *log.Logger
  infoLog        *log.Logger
  warnLog        *log.Logger
  errorLog       *log.Logger
  fatalLog       *log.Logger
}

func NewContext()( *Context ) {
  return &Context{
    NewArguments(),
    log.New(os.Stderr, "DEBUG ", log.Lshortfile | log.LstdFlags ),
    log.New(os.Stderr, "INFO ", log.Lshortfile | log.LstdFlags ),
    log.New(os.Stderr, "WARN ", log.Lshortfile | log.LstdFlags ),
    log.New(os.Stderr, "ERROR ", log.Lshortfile | log.LstdFlags),
    log.New(os.Stderr, "FATAL ", log.Lshortfile | log.LstdFlags),
  }
}


func (ctx *Context) LogDebug(  v ...interface{} ) {
  if *ctx.arguments.verbose {
    ctx.debugLog.Println( v... )
  }
}

func (ctx *Context ) LogInfo( v ...interface{} ) {
  ctx.infoLog.Println( v... )
}

func (ctx *Context ) LogWarn( v ...interface{} ) {
  ctx.warnLog.Println( v... )
}

func (ctx *Context ) LogError( v ...interface{} ) {
  ctx.errorLog.Println( v... )
}

// logs message and terminates application
func (ctx *Context ) LogFatal( v ...interface{} ) {
  ctx.fatalLog.Fatalln( v... )
}
