create or replace function updated_at_function() 
returns trigger language plpgsql as $fn$ 
	begin 
		new.updated_at = now();
		return new;
	end; 
$fn$;

drop trigger if exists created_at_trigger on users;
drop trigger if exists created_at_trigger on realms;

create trigger created_at_trigger before
update on users for each row execute function updated_at_function();

create trigger created_at_trigger before
update on realms for each row execute function updated_at_function();