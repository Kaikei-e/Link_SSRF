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

    console.log(htmlData);

    if (res.status === 200) {
      alert("User registered");

      setSitePreview(htmlData);
    } else {
      alert("Error registering user");
    }
  };

  return (
    <div>
      <h1>Home</h1>
      <h3>Hello </h3>
      <form onSubmit={handleSubmit}>
        <div className="formStyle">
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
      <div>
        <h3>Site Preview</h3>
        <div>
          <p>
            <pre>{sitePreview}</pre>
          </p>
        </div>
      </div>
    </div>
  );
}

export default App;
