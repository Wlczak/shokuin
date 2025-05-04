<!DOCTYPE html>
<html lang="en">
<head>
  <meta charset="UTF-8">
  <meta name="viewport" content="width=device-width, initial-scale=1">
  <title>Coming Soon</title>
  <style>
    html, body {
      height: 100%;
      margin: 0;
      display: flex;
      align-items: center;
      justify-content: center;
      background-color: #000;
    }
    h1 {
      font-family: 'Courier New', Courier, monospace;
      font-size: 6rem;
      color: red;
      text-shadow: 0 0 10px #ff0000, 0 0 20px #ff0000, 0 0 40px #ff0000;
      animation: flicker 1.5s infinite alternate;
    }
    @keyframes flicker {
      0%   { opacity: 1; text-shadow: 0 0 10px #ff0000, 0 0 20px #ff0000, 0 0 40px #ff0000; }
      40%  { opacity: 0.8; text-shadow: 0 0 5px #ff0000, 0 0 10px #ff0000; }
      60%  { opacity: 0.6; text-shadow: 0 0 2px #ff0000; }
      100% { opacity: 1; text-shadow: 0 0 15px #ff0000, 0 0 30px #ff0000, 0 0 60px #ff0000; }
    }
  </style>
</head>
<body>
  <h1>Coming Soon</h1>
</body>
</html>
