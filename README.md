# output of this example of 4 chan

```bash
sort 4chan.txt | sed -e 's/.*msg=//' | uniq -c | sort -n
```

will give next ourput
```bash
   1 3 ^ 5: 243
   3 got message from 4chan
   3 starting 3chan producer
   9 got message from 3chan
   9 starting 2chan producer
   9 starting 3chan consumer
  27 got message from 2chan
  27 starting 1chan producer
  27 starting 2chan consumer
  81 starting 1chan consumer
  81 starting int producer
 243 received int
 243 sending int
 ```