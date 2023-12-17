import type { User } from "./types/user";

export const registerUser = async (username: string, url: string) => {
  const user: User = {
    username: username,
    profile_link: url,
  };

  const endpoint = import.meta.env.VITE_API_ENDPOINT;

  console.log(endpoint);
  console.log(user);
  4173
  const response = await fetch(endpoint + `/register`, {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(user),
  });
  const data = await response.json();
  return data;
};
