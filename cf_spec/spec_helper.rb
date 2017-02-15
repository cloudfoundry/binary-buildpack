require 'machete'
require 'machete/matchers'

`mkdir -p log`
Machete.logger = Machete::Logger.new("log/integration.log")


def skip_if_no_windows_stack
  return if has_windows_stack?

  skip 'cf installation does not have a Windows stack'
end

def has_windows_stack?
  Machete::CF::Stacks.new.execute.include? 'windows2012R2'
end

def diego_enabled?(app_name)
  `cf has-diego-enabled #{app_name}`.chomp == 'true'
end
