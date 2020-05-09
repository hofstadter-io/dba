package schema

#User: @modelset(entity,opencore)
#User: @model()
#User: {
  ID: "uuid" @db("key","key=val","key!val") @required()
  CID: "cuid" @db()
  SHORT: string // & len(SHORT) == 8

  Email: "email" // TODO, turn "email" into string + regex ~> types.Email

  ...
}


