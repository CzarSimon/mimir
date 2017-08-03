const getTwitterData = (req, res, conn) => {
  console.log("this will get twitter data");
  res.status(200).send("getTwitterData\n");
}

module.exports = {
  getTwitterData
}
