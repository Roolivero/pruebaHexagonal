export declare function cmd(): Promise<void>;
export declare function createNewProject({ workspaceId, name, region, }: {
    workspaceId: string;
    name: string;
    region: string;
}): Promise<{
    id: string;
    idName: string;
}>;
export declare function promptExistingProjectName({ workspaceId }: {
    workspaceId: string;
}): Promise<{
    id: string;
    name: string;
    idName: string;
}>;
export declare function promptBegin(): Promise<void>;
/**
 * Ask the user if they want to link to an existing Astro Studio project.
 * @returns A `Promise` for the user’s answer: `true` if they answer yes, otherwise `false`.
 */
export declare function promptLinkExisting(): Promise<boolean>;
/**
 * Ask the user if they want to link to a new Astro Studio Project.
 * **Exits the process if they answer no.**
 * @returns A `Promise` for the user’s answer: `true` if they answer yes.
 */
export declare function promptLinkNew(): Promise<boolean>;
export declare function promptNewProjectName(): Promise<string>;
export declare function promptNewProjectRegion(): Promise<string>;
