import json
import sys

article_info = json.loads(sys.argv[1])

print("Article info")
print(article_info)


sys.exit(0)
