require 'spec_helper'

describe 'CF Binary Buildpack' do
  after do
    Machete::CF::DeleteApp.new.execute(app)
  end

  describe 'deploying a Windows HWC app' do
    let(:app_name) { 'hwc_app' }

    context 'without a command or Procfile' do
      let(:app) { Machete.deploy_app(app_name, buildpack: 'binary-test-buildpack', stack: 'windows2012R2', start_command: 'null') }

      it 'logs a warning message' do
        skip_if_no_windows_stack

        expect(app).to have_logged("Warning: We detected a Web.config in your app. This probably means that you want to use the hwc-buildpack. If you really want to use the binary-buildpack, you must specify a start command.")
      end
    end
  end
end
