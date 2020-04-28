package schema

#User: {
  ID: "uuid"
  CID: "cuid"
  SHORT: string // & len(SHORT) == 8

  Email: "email" // TODO, turn "email" into string + regex ~> types.Email

  ...
}
