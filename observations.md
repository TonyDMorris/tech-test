# observations

## S3 as a database
Considering that the needs of this project are to simply to store json files via a GUID and retreive them the same way, S3 is actually great choice for the POC as it is significantly cheaper than the alternatives.
In fact I have read that orders older than 12 months on amazon are actually stored in S3. 
Should the need to SQL functionality arise AWS athena could be used to query the infromation in the future.

## changes

i have begun to implement some changes to the poc , firstly i changes the cloud formation templates to make the stack region agnostic .

secondly I implemented a new lambda function which I was going to use for storing and retreiving questionaires from an s3 bucket

unfortunatley once i created the cloud formation template and implemented it it seems to have broken the build for one of the other templates.

I have spent some time trying to debug this but I am not really massively experienced with cloud formation templates and was unable to solve it in the time I alloted for the issue.

The Intention was to use ListenQuestionaireCaptureLambda to implement questionaire creation and retreival via the lambda, as the templates are unbuildable i have just written the code as though it was functional so you can see what i intended to do.

I really would have like to implement the changes to the front end to create a taste of questionnaire creation and link up the lambda to retreive the latest questionnaire from s3 whenever the site is loaded but i simply ran out of time .







