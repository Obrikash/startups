"use server";

import { parseServerActionResponse } from "@/lib/utils";
import slugify from "slugify";

export const createPitch = async (
  state: any,
  form: FormData,
  pitch: string,
) => {
  const { title, description, category, link } = Object.fromEntries(
    Array.from(form).filter(([key]) => key !== "pitch"),
  );

  const slug = slugify(title as string, { lower: true, strict: true });

  try {
    const startup = {
      title,
      description,
      category,
      image: link, 
      slug: {
        _type: slug,
        current: slug,
      },
      author: {
        id: 1, 
      },
      pitch,
    };

    
    const response = await fetch("http://localhost:4000/api/startups", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(startup),
    });

    if (!response.ok) {
      const errorData = await response.json();
      throw new Error(errorData.message || "Failed to create pitch");
    }

    const result = await response.json();

    return parseServerActionResponse({
      ...result,
      error: "",
      status: "SUCCESS",
    });
  } catch (error) {
    console.error(error);

    return parseServerActionResponse({
      error: error instanceof Error ? error.message : JSON.stringify(error),
      status: "ERROR",
    });
  }
};