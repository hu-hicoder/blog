import * as React from "react";
import { useState, useEffect } from "react";

type Article = {
  ID: number;
  CreatedAt: string;
  UpdatedAt: string;
  DeletedAt: string | null;
  Title: string;
  Author: string;
  Content: string;
};

const HelloWorld: React.FC = () => {
  const [articles, setArticles] = useState<Article[] | null>(null);

  useEffect(() => {
    const load = async () => {
      await fetch("http://localhost:3001/article")
        .then((res) => res.json())
        .then((data) => setArticles(data["Articles"]))
        .catch((err) => new Error("Cannot Fetch" + err));
    }
    load()
  } , []);
  return (
    <>
      <h1 className="hello">Hello World!!</h1>{" "}
      {articles && articles.map(item => <div key={item.ID}>Author: {item.Author} / Content: {item.Content}</div>)}
    </>
  );
};

export default HelloWorld;
