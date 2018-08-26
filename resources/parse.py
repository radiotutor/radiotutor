#!/usr/bin/python3

import sys
import re

path = sys.argv[1]
the_file = open(path,'r', encoding="utf-8")
lines = the_file.read()

fmt = re.sub("([\r\n].*?)(?:=?\r|\n)(.*?(?:\$CATEGORY: .*))", "", lines, flags=re.MULTILINE)
fmt = re.sub("(?i)<img[^>]*>", "IMG_NEEDED", fmt, flags=re.MULTILINE)
fmt = re.sub("<[^>]*>", "", fmt, flags=re.MULTILINE)
fmt = fmt.replace("\\", "")
fmt = fmt.replace("\"", "\\\"")
fmt = fmt.replace("{", "\",\n  \"answers\": [")
fmt = fmt.replace("}", "  ]\n},")
fmt = re.sub("  name:.*", ",", fmt, flags=re.MULTILINE)
fmt = fmt.replace("// question: ", "{\n  \"number\":")
fmt = fmt.replace("[html]", "")
fmt = fmt.replace("[moodle]", "")
#fmt = fmt.replace("<p>", "")
#fmt = fmt.replace("</p>", "")
	  
fmt = fmt.replace("\t~", "    {\"correct\":false, \"answer\":\"")
fmt = fmt.replace("\t=", "    {\"correct\":true, \"answer\":\"")
fmt = re.sub("^.*(?=::([^\s]+)\s)::([^\s]+)\s", "  \"question\":\"", fmt, flags=re.MULTILINE)
fmt = re.sub("^(?:[\t ]*(?:\r?\n|\r))+", "", fmt, flags=re.MULTILINE)
fmt = re.sub("(	####.*)", "", fmt, flags=re.MULTILINE)
fmt = fmt.replace("#", "")
#fmt = re.sub("<.*>", "THE_HTML", fmt, flags=re.MULTILINE)
fmt = fmt.replace("\t", " ")

lineitt = fmt.splitlines()

x = 0
i = 0
print("[")
while x < len(lineitt):
  if lineitt[x][:5] == "    {":
    i = i + 1
    lineitt[x] = lineitt[x] + "\""
    if i%4 == 0:
      lineitt[x] = lineitt[x] + "}"
    else:
      lineitt[x] = lineitt[x] + "},"

  if x == len(lineitt) - 1:
    lineitt[x] = "}"

  print(lineitt[x])
  x = x + 1
print("]")
