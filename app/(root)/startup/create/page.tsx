import StartupForm from "@/components/StartupForm";
import { auth } from "@/auth";
import { redirect } from "next/navigation";

const Page = async () => {
  return (
    <>
      <section className="pink_container !min-h-[230px]">
        <h1 className="heading">Создай свой Стартап!</h1>
      </section>

      <StartupForm />
    </>
  );
};

export default Page;
