const Login = () => {
    let state = {
        username: '',
        password: ''
    }

    const updateUsernameValue = (e) => {
        state.username = e.target.value
    }

    const updatePasswordValue = (e) => {
        state.password = e.target.value
    }

    const handleLogin = async () => {
        const requestOptions = {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({
                username: state.username,
                password: state.password
            }),
        };

        const response = await fetch(`http://localhost:3001/auth/login`, requestOptions);
        await response.json().then((res) => {
            console.log(res)
            if (res.code == 400) {
                alert(res.message)
            } else {
                alert("logged in!")
                window.localStorage.setItem("token", res.token)
            }
        });
    }

    return (
        <>
            Username
            <input id={"username"} name={"username"} onChange={(e) => updateUsernameValue(e)}/>
            <br/>
            Password
            <input id={"password"} type="password" onChange={(e) => updatePasswordValue(e)} />
            <button onClick={() => handleLogin()}>Login</button>
        </>
    );
}

export default Login;
