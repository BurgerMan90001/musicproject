function Card({ title }: { title: string }) {
  return <div className="flex  bg-black-20">{title}</div>;
}
function Legal() {
  return (
    <>
      <Card title="Privacy Policy" />
      <Card title="Terms of Service" />
    </>
  );
}

export default Legal;
