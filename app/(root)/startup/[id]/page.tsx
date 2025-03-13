import { formatDate } from "@/lib/utils";
import Image from "next/image";
import Link from "next/link";
import { notFound } from "next/navigation";
import React from "react";

import markdownit from "markdown-it";
import View from "@/components/View";

const md = markdownit();

const Page = async ({ params }: { params: Promise<{ id: number }> }) => {
  const id = (await params).id;

  const response = await fetch("http://localhost:4000/api/startups/" + id);
  const { startup: post } = await response.json();
  console.log(post);
  if (!post) return notFound();

  const parsedContent = md.render(post?.pitch || "");
  return (
    <>
      <section className="pink_container pattern !min-h-[230px]">
        <p className="tag">{formatDate(post?.created_at)}</p>
        <h1 className="heading">{post.title}</h1>
        <p className="sub-heading !max-w-5xl">{post.description}</p>
      </section>

      <section className="section_container">
        <img
          src={post.image_url}
          alt="thumbnail"
          className="w-full md:max-w-xl lg:max-w-2xl h-auto rounded-xl"
        ></img>

        <div className="space-y-5 mt-10 max-w-4xl mx-auto">
          <div className="flex-between gap-5">
            <Link
              href={`/user/${post.author?.id}`}
              className="flex gap-2 items-center mb-3"
            >
              <Image
                src={post.author.image_url}
                alt="avatar"
                width={64}
                height={64}
                className="rounded-full drop-shadow-lg"
              />

              <div>
                <p className="text-20-medium">{post.author.name}</p>
                <p className="text-16-medium !text-black-300">
                  @{post.author.username}
                </p>
              </div>
            </Link>
            <p className="category-tag">{post.category}</p>
          </div>
          <h3 className="text-30-bold">Детали Стартапа</h3>
          {parsedContent ? (
            <article
              dangerouslySetInnerHTML={{ __html: parsedContent }}
              className="prose max-w-4xl font-work-sans break-all"
            ></article>
          ) : (
            <p className="no-result">No details provided</p>
          )}
        </div>
        <hr className="divider" />

        <View views={post.views} />
      </section>
    </>
  );
};

export default Page;
