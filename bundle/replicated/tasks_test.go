package replicated

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTasksParse(t *testing.T) {
	tasksText := `ID                                  DESCRIPTION                                                             STARTED                    STATUS
0f09e0088843fc20104eb3d1717a6d44    Docker commands                                                         2017-01-03 22:26:04 UTC    Executing
1f44063b6d166f555c4e4f2cbc15b8ee    Request scheduler replicated support bundle                             2017-01-03 22:26:04 UTC    Executing
3413fe7608999d33c4db8dcceaf42a3b    Replicated daemon files                                                 2017-01-03 22:26:04 UTC    Executing
7ab92ca89df262d58f9413a8d1a408eb    Replicated docker commands                                              2017-01-03 22:26:04 UTC    Executing
8128191403a5b83c2bd148b235fbd3bf    Ledis db dump                                                           2017-01-03 22:26:04 UTC    Executing
81a37ce58dae040019dec5ab0522c229    Journald Logs                                                           2017-01-03 22:26:04 UTC    Executing
92b9c041dd66964a3ac1b4b00350f8bb    Pack support bundle                                                     2017-01-03 22:26:04 UTC    Queued
95c849dbdacef219bf0d3d316084d59b    Scheduled task to expunge unused metrics data                           2017-01-03 22:16:57 UTC    Sleeping
9a4a554abd19922b738dd5adc6fe20d1    Task list                                                               2017-01-03 22:26:04 UTC    Executing
d3b02e9e3f1f38b51d0c696fc78d4510    Replicated daemon system commands                                       2017-01-03 22:26:04 UTC    Executing
e2e95bf2912bbc8fc2b564bdbd4c9160    Scheduled task to expunge app snapshots greater than "max_snapshots"    2017-01-03 22:16:57 UTC    Sleeping
e3740ddf51079f3869b51150c57f4faf    Audit Log                                                               2017-01-03 22:26:04 UTC    Executing
e71b04a5e121439eab24a06f3f511fe7    Cleaning up piper containers                                            2017-01-03 22:16:57 UTC    Sleeping
f08646caaf6142ac931c9271f6d7deed    Scheduled license sync                                                  2017-01-03 22:16:57 UTC    Sleeping
f3eb870d71e78431051622b74457062e    Scheduled task to expunge support bundles older than an hour            2017-01-03 22:16:57 UTC    Sleeping
fe45e5da36e34f48a2ccdeae28b9f67f    Scheduled check for app upgrades                                        2017-01-03 20:16:57 UTC    Sleeping
`
	tasks, err := ParseTasks(tasksText)
	assert.Nil(t, err)
	assert.Equal(t, "2h9m7s", tasks.TimeRange.String())
}
