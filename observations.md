# observations
&nbsp;
&nbsp;

##  Can we build systems that avoid databases?  Previously scaling databases has been a problem, so can we build systems that don't use them?  What trade-offs does this give?

&nbsp;

In my opinion I would say that S3 actually is a database , not a traditional one but a database none the less but
considering the needs of this project are to simply store json files via a GUID and retrieve them the same way, S3 is actually a great choice for the POC as it is significantly cheaper than the alternatives.

In fact I have read that orders older than 12 months on amazon are actually stored in S3. 

Should the need to SQL functionality arise AWS athena could be used to query the information in the future.

My second thought on this would be that if the project does become more complex and the capabilities of a traditional database are required you might find that the work involved to migrate this project over might be considerably more than the upfront expense of just using a managed service like RDS in the first place , although this is only my opinion.

To answer the question i would say no we can't really avoid databases , we can use other storage methods such as S3 but I would just consider this to be using the chosen service as a database and not an actual replacement, but I am happy to be convinced otherwise.

&nbsp;


##  Is this tool actually useful?  Will our team love it?  Will it give us interesting actionable data?  If so can we incorporate it into our SASS product?

&nbsp;

I think the tool provides a good underpinning for something that could be useful.

If this poc where coupled with an automated e-mail campaign to distribute the weekly Questionnaires and something was written to perform some basic data analysis on the weekly/monthly results , I feel this could be a useful metric but it would entirley depend on how hard you push to have the Questionnaires completed by as many people as possible and the level of analysis undertaken.

Will your team love it? well that entirely depends on the individuals in your team but from personal experience I have completed probably 100s of these types of surveys and i can't say that I have enjoyed any of them in my honest opinion.


&nbsp;


##  Identify one aspect you would like to improve

&nbsp;


I have begun to implement some changes to the poc , firstly i changed the cloud formation templates to make the stack region-agnostic as it was hard coded to eu-west-1.

secondly I implemented a new lambda function which I was going to use for storing and retrieving questionnaires from an s3 bucket

unfortunately once i created the cloud formation template and implemented it it seems to have broken the build for one of the other templates.

I have spent some time trying to debug this but I am not really massively experienced with cloud formation templates and was unable to solve it in the time I alloted for the issue.

The Intention was to use ListenQuestionnaireCaptureLambda to implement questionnaire creation and retrieval via the lambda, as the templates are not buildable i have just written the code as though it was functional so you can see what i intended to do.

I really would have like to implement the changes to the front end to create a taste of questionnaire creation and link up the lambda to retrieve the latest questionnaire from s3 whenever the site is loaded but i simply ran out of time .


&nbsp;


## Critique and suggest/action improvements



&nbsp;

I think the basic skeleton of the product is there , I'm not really sure what to put in this section so i will just list some things I would start on if I where given this to work on.

* implement a data storage interface and an S3 concrete implementation so that we are not tied to S3 in the future and so that our code is easily testable.

* configure the cloud formation templates to generate an api that exposes routes for Prod Dev and Test environments along with corresponding buckets.

* implement amazon cognito for accessing the api gateway or write the logic to have this sit behind your existing infrastructure
and take advantage of any authentication and authorization you have up and running already.

* rebuild the project into golang sub-modules so that structs and logic can be shared across the individual lambdas.

* implement configuration structs so we can bring in environment variables in a nice clear way.

* implement an application struct for each package which would be used to build each lambdas functionality via composition.

* lots of tests

* implement some sort of pipeline to push future changes to production via release tags as opposed to uploading zip files.

## conclusion

I have enjoyed understanding the project a lot especially the cloud formation templates,  I am just sorry that I have not got much more time to work on this due to my current commitments.

Thanks for your time all the best.

Tony.












