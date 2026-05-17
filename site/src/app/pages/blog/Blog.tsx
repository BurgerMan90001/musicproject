import { Outlet } from "react-router";
import BlogFooter from "./BlogFooter";
import Header from "../../../components/header/Header";

// const Content = ({
//   title,
//   author,
//   content,
// }: {
//   title: string;
//   author: string;
//   content: string;
// }) => {
//   return (
//     <main className="">
//       <h1>{title}</h1>
//       <h2 className="color-text-subtle">{author}</h2>
//       <div>{content}</div>
//     </main>
//   );
// };

const Blog = () => {
  return (
    <div>
      <Header />
      <Outlet />
      <BlogFooter />
    </div>
  );
};

export default Blog;
