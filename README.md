# ent-join
entでjoinをテストする

`uploaded_content`テーブル
|filename|
|----|
|sample1.mp4|
|sample2.mp4|

`content`テーブル
|filename|uploaded_content_filename|
|----|----|
|content1_sample1.mp4|sample1.mp4|
|content2_sample1.mp4|sample1.mp4|
|content1_sample2.mp4|sample2.mp4|
|content2_sample2.mp4|sample2.mp4|
|content3_sample1.mp4|sample1.mp4|

`content_movie_metadata`テーブル
|filename|width|height|
|----|----|----|
|content1_sample1.mp4|1920|1080|
|content1_sample2.mp4|1920|1080|
|content2_sample1.mp4|1920|1080|
|content2_sample2.mp4|1920|1080|
|content3_sample1.mp4|1080|720|
