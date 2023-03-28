<script lang='ts'>
    import { Login } from "../api/login";
    import { goto } from "$app/navigation";
    let username = '';
    let password = '';
    let message = '';

    function LoginUser() {
        console.log("Attempting to log in user")
        Login(username, password).then(
            response => {
                message = "Succesful Login"
                goto('/dadgpt/'.concat(response["UserId"]))
            }
        ).catch(error => {
            message = "Unsuccesful Login"
            alert(message)
        })
    }

</script>

<div>
    <form on:submit|preventDefault={() => LoginUser()}>
        <p>{message}</p>
        <input type="text" name="username" placeholder="Username" bind:value={username}>
        <input type="password" name="password" placeholder="Password" bind:value={password}>
        <button type="submit">Login</button>
    </form>

</div>
