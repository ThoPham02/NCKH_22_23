import client from "../../apis";
import "./login.css";

const Login = () => {
  const handleLoginSubmit = (e) => {
    e.preventDefault();
    
    client.post('/user/login', {
      username: e.target.username.value,
      password: e.target.password.value
    })
      .then(response => console.log(response.data))
      .catch(error => console.error(error));

  };

  return (
    <div id="login">
      <div className="login-heading">
        <h1>Đăng nhập tài khoản</h1>
      </div>
      <form onSubmit={handleLoginSubmit} >
        <div className="input">
          <h2>Tên đăng nhập</h2>
          <input type="text" name="username" placeholder="Username" />
        </div>
        <div className="input">
          <h2>Mật khẩu</h2>
          <input type="password" name="password" placeholder="Password" />
        </div>
        <button
          id="btn"
          className="btn"
          type="submit"
        >
          Đăng nhập
        </button>
      </form>
    </div>
  );
};

export default Login;
