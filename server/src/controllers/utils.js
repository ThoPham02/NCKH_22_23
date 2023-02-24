module.exports = {
    handleResponse: function(res, err, data) {
        if (err) {
            console.log(err)
            res.json({
              status: 400,
              message: "Bad request",
            });
            return;
          }
          res.json({
            status: 200,
            message: "OK",
            data: data,
          });
    }
}