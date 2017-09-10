# R server for preparing data, querying data, prediction services.
library(Rserve)

#Prepare training and testing dataset here.
prepareData <- function ()
{
# create training and testing data from data files

#load the state wise data
APdata <- read.csv(file="/home/shajulin/shaju/workspace/Eclipseworkspace/go/src/github.com/RevAirQuality/data-arunachal.csv",sep=",",head=TRUE)
Pondydata <- read.csv(file="/home/shajulin/shaju/workspace/Eclipseworkspace/go/src/github.com/RevAirQuality/data-puducherry.csv",sep=",",head=TRUE)
TNdata <- read.csv(file="/home/shajulin/shaju/workspace/Eclipseworkspace/go/src/github.com/RevAirQuality/data-tamilnadu.csv",sep=",",head=TRUE)
Karnatakadata <- read.csv(file="/home/shajulin/shaju/workspace/Eclipseworkspace/go/src/github.com/RevAirQuality/data-karnataka.csv",sep=",",head=TRUE)
Keraladata <- read.csv(file="/home/shajulin/shaju/workspace/Eclipseworkspace/go/src/github.com/RevAirQuality/data-kerala.csv", sep=",", head=TRUE)
#Assamdata <- read.csv(file="/home/shajulin/shaju/workspace/Eclipseworkspace/go/src/github.com/RevAirQuality/data-assam.csv",sep=",",head=TRUE)


# raw data
rawData = rbind(APdata, Pondydata, TNdata, Karnatakadata, Keraladata)
nrow(rawData)

# convert date to numeric
rawData$dateNum <- as.numeric(rawData$Sampling.Date)

# combine filtered data
combinedData <- rawData[which(rawData$SO2 != 'NA' & rawData$NO2 != 'NA' & rawData$RSPM.PM10 != 'NA' & rawData$Stn.Code != 'NA'),]
nrow(combinedData)

# create a unique station time
maxTime <- max(combinedData$dateNum)
combinedData$OrderTime <- seq.int(nrow(combinedData))
combinedData$UniqueStationTime <- (combinedData$Stn.Code + combinedData$OrderTime + combinedData$dateNum)

#order NO2 value
OrderNO2Data <- combinedData[order(combinedData[,11], combinedData[,7], decreasing=FALSE),]
OrderNO2Data$SNo <- seq.int(nrow(OrderNO2Data))
nrow(OrderNO2Data)
head(OrderNO2Data)


#splitdf function will return a list of training and testing sets
splitdf <- function(OrderNO2Data, seed=NULL) {
if (!is.null(seed)) set.seed(seed)
index <- 1:nrow(OrderNO2Data)
trainindex <- sample(index, trunc(length(index)/2))
trainset <- OrderNO2Data[trainindex, ]
testset <- OrderNO2Data[-trainindex, ]
list(trainset=trainset,testset=testset)
}

#apply the function
splits <- splitdf(OrderNO2Data, seed=808)

#it returns a list - two data frames called trainset and testset
str(splits)

# there are 50-50 observations foreach data frame
lapply(splits,nrow)

#view the first few columns in each data frame
lapply(splits,head)

# save the training and testing sets as data frames
training <- splits$trainset
testing <- splits$testset

# write the values to the training and testing dataset files
write.table(training, "/home/shajulin/shaju/tempfolder/training.csv", col.names=TRUE, sep=",")
write.table(testing, "/home/shajulin/shaju/tempfolder/testing.csv", col.names=TRUE, sep=",")
}

# QueryData with number of frequency here.
queryData <- function ()
{
  testing <- read.csv(file="/home/shajulin/shaju/tempfolder/testing.csv",sep=",",head=TRUE)
  table(unlist(testing$Location.of.Monitoring.Station))
  write.csv(table(unlist(testing$Location.of.Monitoring.Station)), file = "/home/shajulin/shaju/tempfolder/RevenueCities.csv")

}

# Revenue calculation with number of frequency here.
revenueData <- function ()
{

  revenueData <- read.csv(file="/home/shajulin/shaju/tempfolder/RevenueCities.csv",sep=",",head=TRUE)
  rows <- nrow(revenueData)
  for (i in 1:rows){
    if(revenueData$Freq < 50) {
      revenueData$Credit <- revenueData$Freq * 3 * 50
    }
    else {
      revenueData$Credit <- revenueData$Freq * 3 * 63.3
    }
  }
  write.csv(revenueData, file="/home/shajulin/shaju/tempfolder/RevenueCitiesFinal.csv")

}

# predictSO2 calculation here
predictSO2 <- function()
{
  training <- read.csv(file="/home/shajulin/shaju/tempfolder/training.csv",sep=",",head=TRUE)
  testing <- read.csv(file="/home/shajulin/shaju/tempfolder/testing.csv",sep=",",head=TRUE)

  ################### RFM Prediction Approach
  library(randomForest)

  modelSO2 <- randomForest(SO2 ~ NO2 + RSPM.PM10 +dateNum + Stn.Code + SNo, data = training, importance=TRUE, keep.forest=TRUE, ntree=100, mtry=2)


  #predict the outcome of the testing data
  predictedSO2 <- predict(modelSO2, testing, type="response", predict.all=FALSE, proximity=FALSE, nodes=FALSE)

  ############################ Find Error for the prediction
  actualSO2 <- testing$SO2
  rsqSO2 <- 1-sum((actualSO2 - predictedSO2)^2)/sum((actualSO2-mean(actualSO2))^2)
  print(rsqSO2)

  x=rnorm(100)
  y=rnorm(100,5,1)
  plot(training$UniqueStationTime, training$SO2,col="blue",xlab="(MeasurementTime + Location) Unit ", ylab="SO2 value ", cex.lab=1.1)
  lines(testing$UniqueStationTime, predictedSO2, col="red")


  # Plot the prediction graph
  SO2Filename <- tempfile("SO2predict", fileext = ".pdf")
  tempdir()
  png(SO2Filename, width = 5, height = 4, units = 'in', res = 300)
  x=rnorm(100)
  y=rnorm(100,5,1)
  plot(training$UniqueStationTime, training$SO2,col="blue",xlab="(Measured Date + Location) Unit ", ylab="SO2 value ", cex.lab=1.1)
  grid (NULL,NULL, lty = 6, col = "cornsilk2")
  lines(testing$UniqueStationTime, predictedSO2, col="red")
  dev.off()

  # Plot the prediction graph to local disk
  png("/home/shajulin/shaju/tempfolder/SO2predict.pdf",width=5,height=4, units = 'in', res = 300)
   x=rnorm(100)
   y=rnorm(100,5,1)
   plot(training$UniqueStationTime, training$SO2,col="blue",xlab="(Measured Date + Location) Unit ", ylab="SO2 value ", cex.lab=1.1)
   grid (NULL,NULL, lty = 6, col = "cornsilk2")
   lines(testing$UniqueStationTime, predictedSO2, col="red")
   dev.off()

}


# predictNO2 calculation here
predictNO2 <- function()
{

################### RFM Prediction Approach
library(randomForest)

modelNO2 <- randomForest(NO2 ~ SO2 + RSPM.PM10 +dateNum + Stn.Code + SNo, data = training, importance=TRUE, keep.forest=TRUE, ntree=100, mtry=2)

#predict the outcome of the testing data
predictedNO2 <- predict(modelNO2, testing, type="response", predict.all=FALSE, proximity=FALSE, nodes=FALSE)

############################ Find Error for the prediction
actualNO2 <- testing$NO2
rsqNO2 <- 1-sum((actualNO2 - predictedNO2)^2)/sum((actualNO2-mean(actualNO2))^2)
print(rsqNO2)

x=rnorm(100)
y=rnorm(100,5,1)
plot(training$UniqueStationTime, training$NO2,col="blue",xlab="(MeasurementTime + Location) Unit ", ylab="NO2 value ", cex.lab=1.1)
lines(testing$UniqueStationTime, predictedNO2, col="red")


# Plot the prediction graph
NO2Filename <- tempfile("NO2predict", fileext = ".pdf")
tempdir()
png(NO2Filename, width = 5, height = 4, units = 'in', res = 300)
x=rnorm(100)
y=rnorm(100,5,1)
plot(training$UniqueStationTime, training$NO2,col="blue",xlab="(Measured Date + Location) Unit ", ylab="NO2 value ", cex.lab=1.1)
grid (NULL,NULL, lty = 6, col = "cornsilk2")
lines(testing$UniqueStationTime, predictedNO2, col="red")
dev.off()

# Plot the prediction graph to local disk
png("/home/shajulin/shaju/tempfolder/RevAirQuality/NO2predict.pdf",width=5,height=4, units = 'in', res = 300)
 x=rnorm(100)
 y=rnorm(100,5,1)
 plot(training$UniqueStationTime, training$NO2,col="blue",xlab="(Measured Date + Location) Unit ", ylab="NO2 value ", cex.lab=1.1)
 grid (NULL,NULL, lty = 6, col = "cornsilk2")
 lines(testing$UniqueStationTime, predictedNO2, col="red")
 dev.off()

}

# predictRSPM calculation here
predictRSPM <- function()
{


################### RFM Prediction Approach
library(randomForest)

modelRSPM <- randomForest(RSPM.PM10 ~ NO2 + SO2 +dateNum + Stn.Code + SNo, data = training, importance=TRUE, keep.forest=TRUE, ntree=100, mtry=2)


#predict the outcome of the testing data
predictedRSPM <- predict(modelRSPM, testing, type="response", predict.all=FALSE, proximity=FALSE, nodes=FALSE)

############################ Find Error for the prediction
actualRSPM <- testing$RSPM.PM10
rsqRSPM <- 1-sum((actualRSPM - predictedRSPM)^2)/sum((actualRSPM-mean(actualRSPM))^2)
print(rsqRSPM)

x=rnorm(100)
y=rnorm(100,5,1)
plot(training$UniqueStationTime, training$RSPM.PM10,col="blue",xlab="(MeasurementTime + Location) Unit ", ylab="RSPM value ", cex.lab=1.1)
lines(testing$UniqueStationTime, predictedRSPM, col="red")


# Plot the prediction graph
RSPMFilename <- tempfile("RSPMpredict", fileext = ".pdf")
tempdir()
png(RSPMFilename, width = 5, height = 4, units = 'in', res = 300)
x=rnorm(100)
y=rnorm(100,5,1)
plot(training$UniqueStationTime, training$RSPM.PM10,col="blue",xlab="(Measured Date + Location) Unit ", ylab="RSPM value ", cex.lab=1.1)
grid (NULL,NULL, lty = 6, col = "cornsilk2")
lines(testing$UniqueStationTime, predictedRSPM, col="red")
dev.off()

# Plot the prediction graph to local disk
png("/home/shajulin/shaju/tempfolder/RevAirQuality/RSPMpredict.pdf",width=5,height=4, units = 'in', res = 300)
 x=rnorm(100)
 y=rnorm(100,5,1)
 plot(training$UniqueStationTime, training$RSPM.PM10,col="blue",xlab="(Measured Date + Location) Unit ", ylab="RSPM value ", cex.lab=1.1)
 grid (NULL,NULL, lty = 6, col = "cornsilk2")
 lines(testing$UniqueStationTime, predictedRSPM, col="red")
 dev.off()

}

#Start the server now
run.Rserve()
