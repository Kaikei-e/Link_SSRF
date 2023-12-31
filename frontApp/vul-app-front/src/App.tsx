import "./App.css";
import React from "react";
import { registerUser } from "./lib/registerUser";

function App() {
  const [username, setUsername] = React.useState("");
  const [url, setUrl] = React.useState("");
  const [sitePreview, setSitePreview] = React.useState("");

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault();

    const res = await registerUser(username, url);
    const htmlData = res.message;

    if (res.status === 200) {
      alert("User registered");

      setSitePreview(htmlData);
    } else {
      alert("Error registering user");
    }
  };

  return (
    <div
      style={{
        flexDirection: "column",
        alignItems: "center",
        justifyContent: "center",
        height: "100%",
        width: "100%",
        display: "block",
      }}
    >
      <h1>Home</h1>
      <h3>Hello </h3>
      <h3>Register</h3>
      <form onSubmit={handleSubmit}>
        <div
          style={{
            flexDirection: "column",
            alignItems: "center",
            justifyContent: "center",
            display: "flex",
          }}
        >
          <label htmlFor="username">Name</label>
          <input
            type="text"
            name="username"
            onChange={(e) => setUsername(e.target.value)}
            value={username}
            required
          />
          <label htmlFor="url">URL</label>
          <input
            type="url"
            name="url"
            onChange={(e) => setUrl(e.target.value)}
            value={url}
            required
          />
          <button type="submit">Submit</button>
        </div>
      </form>
      <div
        style={{
          width: "100%",
          flexDirection: "column",
          alignItems: "center",
          justifyContent: "center",
          display: "flex",
        }}
      >
        <h3>Site Preview</h3>
        <div>
          <p
            style={{
              wordBreak: "break-all",
            }}
          >
            {sitePreview}
          </p>
        </div>
      </div>
    </div>
  );
}

export default App;
