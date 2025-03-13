"use client";

import React, { useState, useActionState } from "react";
import { Input } from "@/components/ui/input";
import { Textarea } from "@/components/ui/textarea";
import { Button } from "@/components/ui/button";
import { Send } from "lucide-react";
import { formSchema } from "@/lib/validation";
import { z } from "zod";
import { useToast } from "@/hooks/use-toast";
import { useRouter } from "next/navigation";
import { createPitch } from "@/lib/actions";

const StartupForm = () => {
  const [errors, setErrors] = useState<Record<string, string>>({});
  const [pitch, setPitch] = useState("");
  const { toast } = useToast();
  const router = useRouter();

  const handleFormSubmit = async (prevState: any, formData: FormData) => {
    try {
      const formValues = {
        title: formData.get("title") as string,
        description: formData.get("description") as string,
        category: formData.get("category") as string,
        link: formData.get("link") as string,
        pitch,
      };

      await formSchema.parseAsync(formValues);

      const result = await createPitch(prevState, formData, pitch);

      if (result.status == "SUCCESS") {
        toast({
          title: "Успех",
          description: "Ваш Стартап был успешно создан",
        });

        router.push(`/startup/${result._id}`);
      }

      return result;
    } catch (error) {
      if (error instanceof z.ZodError) {
        const fieldErorrs = error.flatten().fieldErrors;

        setErrors(fieldErorrs as unknown as Record<string, string>);

        toast({
          title: "Ошибка",
          description: "Проверьте значения и попробуйте заново",
          variant: "destructive",
        });

        return { ...prevState, error: "Validation failed", status: "ERROR" };
      }

      toast({
        title: "Error",
        description: "An unexpected error has occurred",
        variant: "destructive",
      });

      return {
        ...prevState,
        error: "An unexpected error has occurred",
        status: "ERROR",
      };
    }
  };

  const [state, formAction, isPending] = useActionState(handleFormSubmit, {
    error: "",
    status: "INITIAL",
  });

  return (
    <form action={formAction} className="startup-form">
      <div>
        <label htmlFor="title" className="startup-form_label">
          Название
        </label>
        <Input
          id="title"
          name="title"
          className="startup-form_input"
          required
          placeholder="Твой Стартап..."
        />

        {errors.title && <p className="startup-form_error">{errors.title}</p>}
      </div>

      <div>
        <label htmlFor="description" className="startup-form_label">
          Описание
        </label>
        <Textarea
          id="description"
          name="description"
          className="startup-form_textarea"
          required
          placeholder="Что же это может быть..."
        />

        {errors.description && (
          <p className="startup-form_error">{errors.description}</p>
        )}
      </div>

      <div>
        <label htmlFor="category" className="startup-form_label">
          Категория
        </label>
        <Input
          id="category"
          name="category"
          className="startup-form_input"
          required
          placeholder="Категория (Техника, Здоровье, Образование...)"
        />

        {errors.category && (
          <p className="startup-form_error">{errors.category}</p>
        )}
      </div>

      <div>
        <label htmlFor="link" className="startup-form_label">
          Ссылка на изображение
        </label>
        <Input
          id="link"
          name="link"
          className="startup-form_input"
          required
          placeholder="URL ссылка"
        />

        {errors.link && <p className="startup-form_error">{errors.link}</p>}
      </div>

      <div>
        <label htmlFor="pitch" className="startup-form_label">
          Детали Стартапа
        </label>
        <Input
          id="pitch"
          name="pitch"
          className="startup-form_input"
          required
          placeholder="Какие же детали..."
        />

        {errors.link && <p className="startup-form_error">{errors.link}</p>}
      </div>

      <Button
        type="submit"
        className="startup-form_btn text-white"
        disabled={isPending}
      >
        {isPending ? "Подтверждение..." : "Выложи свой Стартап"}
        <Send className="size-6 ml-2" />
      </Button>
    </form>
  );
};

export default StartupForm;
