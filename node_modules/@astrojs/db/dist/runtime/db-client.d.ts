import type { LibSQLDatabase } from 'drizzle-orm/libsql';
import { type SqliteRemoteDatabase } from 'drizzle-orm/sqlite-proxy';
type LocalDbClientOptions = {
    dbUrl: string;
    enableTransations: boolean;
};
export declare function createLocalDatabaseClient(options: LocalDbClientOptions): LibSQLDatabase;
type RemoteDbClientOptions = {
    dbType: 'studio' | 'libsql';
    appToken: string;
    remoteUrl: string | URL;
};
export declare function createRemoteDatabaseClient(options: RemoteDbClientOptions): SqliteRemoteDatabase<Record<string, never>> | LibSQLDatabase<Record<string, never>>;
export {};
