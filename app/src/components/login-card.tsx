import type { ComponentProps, FormEvent, ReactNode } from "react";

import { Button } from "@/components/ui/button";
import {
  Card,
  CardContent,
  CardDescription,
  CardHeader,
  CardTitle,
} from "@/components/ui/card";
import {
  Field,
  FieldGroup,
  FieldLabel,
  FieldSeparator,
} from "@/components/ui/field";
import { Input } from "@/components/ui/input";
import { cn } from "@/lib/utils";

export type SocialOption = {
  label: string;
  redirect: string;
  icon?: ReactNode;
};

type LoginCardProps = Omit<ComponentProps<"div">, "title" | "onSubmit"> & {
  title: ReactNode;
  description?: ReactNode;
  magicLink?: boolean;
  social?: SocialOption[];
  onMagicLink?: (email: string) => void;
};

export function LoginCard({
  className,
  title,
  description,
  magicLink = true,
  social = [],
  onMagicLink,
  ...props
}: LoginCardProps) {
  const hasSocial = social.length > 0;
  const showSeparator = hasSocial && magicLink;

  const handleSubmit = (event: FormEvent<HTMLFormElement>) => {
    event.preventDefault();
    if (!onMagicLink) return;
    const data = new FormData(event.currentTarget);
    onMagicLink(String(data.get("email") ?? ""));
  };

  return (
    <Card className={cn("w-full max-w-sm", className)} {...props}>
      <CardHeader className="text-center">
        <CardTitle className="text-xl">{title}</CardTitle>
        {description ? <CardDescription>{description}</CardDescription> : null}
      </CardHeader>
      <CardContent>
        <form onSubmit={handleSubmit}>
          <FieldGroup>
            {hasSocial ? (
              <Field>
                {social.map(({ label, redirect, icon }) => (
                  <Button key={label} variant="outline" asChild>
                    <a href={redirect}>
                      {icon}
                      Continue with {label}
                    </a>
                  </Button>
                ))}
              </Field>
            ) : null}

            {showSeparator ? (
              <FieldSeparator className="*:data-[slot=field-separator-content]:bg-card">
                Or continue with
              </FieldSeparator>
            ) : null}

            {magicLink ? (
              <>
                <Field>
                  <FieldLabel htmlFor="email">Email</FieldLabel>
                  <Input
                    id="email"
                    name="email"
                    type="email"
                    placeholder="m@example.com"
                    required
                  />
                </Field>
                <Field>
                  <Button type="submit">Send me a magic link</Button>
                </Field>
              </>
            ) : null}
          </FieldGroup>
        </form>
      </CardContent>
    </Card>
  );
}
