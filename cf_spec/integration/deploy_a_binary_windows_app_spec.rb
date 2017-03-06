require 'spec_helper'

describe 'CF Binary Buildpack' do
  before(:all) { skip_if_no_windows_stack }
  let(:buildpack) { ENV.fetch('SHARED_HOST')=='true' ? 'binary_buildpack' : 'binary-test-buildpack' }

  after do
    Machete::CF::DeleteApp.new.execute(app)
  end

  describe 'deploying a Windows batch script' do
    let(:app_name) { 'windows_app' }

    context 'when specifying a buildpack' do
      let(:app) { Machete.deploy_app(app_name, buildpack: buildpack, stack: 'windows2012R2') }

      it 'deploys successfully' do
        expect(app).to be_running

        expect(app).to have_logged("Hello, world!")
      end
    end

    context 'without specifying a buildpack' do
      let(:app) { Machete.deploy_app(app_name, stack: 'windows2012R2') }

      it 'fails to stage' do
        expect(app).not_to be_running

        if diego_enabled?(app_name)
          expect(app).to have_logged('None of the buildpacks detected a compatible application')
        else
          expect(app).to have_logged('An app was not successfully detected by any available buildpack')
        end
      end
    end

    context 'without a command or Procfile' do
      let(:app) { Machete.deploy_app(app_name, buildpack: buildpack, stack: 'windows2012R2', start_command: 'null') }

      it 'logs an error message' do
        expect(app).to have_logged("Error: no start command specified during staging or launch")
      end
    end
  end
end
